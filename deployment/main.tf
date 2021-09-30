locals {
  cluster_name = "${var.project_name}-eks-${random_string.suffix.result}"
}

resource "random_string" "suffix" {
  length  = 8
  special = false
}

provider "aws" {
  region = var.aws_region
  shared_credentials_file = var.credential_source
}

provider "helm" {
  alias = "local_helm"
  kubernetes {
    host = data.aws_eks_cluster.cluster.endpoint
    cluster_ca_certificate = base64decode(data.aws_eks_cluster.cluster.certificate_authority.0.data)
    token = data.aws_eks_cluster_auth.cluster.token
  }
}

provider "kubernetes" {
  host                   = data.aws_eks_cluster.cluster.endpoint
  cluster_ca_certificate = base64decode(data.aws_eks_cluster.cluster.certificate_authority.0.data)
  token                  = data.aws_eks_cluster_auth.cluster.token
}

data "aws_availability_zones" "available" {}

module "vpc" {
  source  = "terraform-aws-modules/vpc/aws"
  version = "2.78.0"

  name                 = "${var.project_name}-vpc"
  cidr                 = "10.0.0.0/16" # Sets the starting address range
  azs                  = data.aws_availability_zones.available.names

  # Subnets must be in at least two subnets
  private_subnets      = ["10.0.1.0/24", "10.0.2.0/24", "10.0.3.0/24"]
  public_subnets       = ["10.0.4.0/24", "10.0.5.0/24", "10.0.6.0/24"]

  # AWS VPC will configure routing table for everything to be routed through single NAT gateway, objective is to cut costs: https://circleci.com/blog/getting-started-with-kubernetes-how-to-set-up-your-first-cluster/
  enable_nat_gateway   = true

  # From: https://registry.terraform.io/modules/terraform-aws-modules/vpc/aws/1.41.0
  single_nat_gateway   = true

  # Required by EKS: https://docs.aws.amazon.com/eks/latest/userguide/network_reqs.html
  enable_dns_hostnames = true

  public_subnet_tags = {
    "kubernetes.io/cluster/${local.cluster_name}" = "shared"
    "kubernetes.io/role/elb"                      = "1"
  }

  private_subnet_tags = {
    "kubernetes.io/cluster/${local.cluster_name}" = "shared"
    "kubernetes.io/role/internal-elb"             = "1"
  }
}

# Important: https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/eks_cluster
module "eks" {
  source          = "terraform-aws-modules/eks/aws"

  cluster_name    = local.cluster_name
  cluster_version = "1.19"
  subnets         = module.vpc.private_subnets
  vpc_id = module.vpc.vpc_id

  manage_aws_auth  = false

  # Enable EKS logging: https://github.com/jbrt/aws-eks/blob/master/terraform/1-Create-EKS-Cluster/eks.tf
  cluster_enabled_log_types = ["audit", "authenticator"]

  # Keep things as simple as possible
  # Defaults: https://github.com/terraform-aws-modules/terraform-aws-eks/blob/master/local.tf
  node_groups = {
    eks-nodes = {
      name = "eks-nodes"
      instance_type = "c1.medium" # From autoscaling group instance type options

      desired_capacity = 6
      max_capacity     = 7
      min_capacity     = 2
    }
  }
}

// NOTE: Related to https://github.com/hashicorp/terraform-provider-helm/issues/427
// Must create the namespace explicitly before attempting to deploy resources via helm into the namespace
resource "kubernetes_namespace" "ingress-nginx" {
  metadata {
    name = "ingress-nginx"
    annotations = {
      name = "ingress-nginx"
    }
  }
}

# TODO: Any way to get all chart names automatically?
# Deploy services
variable "services" {
  description = "Service names"
  default = ["api-store", "grpc", "public-api", "web-frontend", "ingress-nginx-controller", "ingress", "external-dns"]
}

variable "relativeChartPath" {
  default = "../charts"
}

resource "helm_release" "services" {
  provider = helm.local_helm

  count = length(var.services)
  name = var.services[count.index]

  # Ensure service account exists before attempted deployment of external-dns
  depends_on = [
    kubernetes_namespace.ingress-nginx,
    kubernetes_service_account.external-dns,
    kubernetes_cluster_role.external-dns,
    kubernetes_cluster_role_binding.external-dns-viewer,
    aws_iam_role.external-dns
  ]

  values = [
    file("../env/charts/values-prod.yaml")
  ]

  chart = "${var.relativeChartPath}/${var.services[count.index]}"
}
