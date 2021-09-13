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

# Deploy services
variable "services" {
  description = "Service names"
  default = ["api-store", "grpc", "public-api", "web-frontend", "ingress", "external-dns"]
}

variable "relativeChartPath" {
  default = "../charts"
}

# TODO: May need to do something similar for ingress

# Create cert-managar namespace and resource before progressing
resource "kubernetes_namespace" "cert-manager" {
  metadata {
    annotations = {
      name = "cert-manager"
    }

    name = "cert-manager"
  }
}

resource "helm_release" "cert-manager" {
  name       = "cert-manager"
  repository = "https://charts.jetstack.io"
  chart      = "cert-manager"
  version    = "v1.4.0"
  namespace  = "cert-manager"
  # namespace  = rancher2_namespace.cert-manager.name
  # lint       = true
  # atomic     = true
  values = [
    file("./values/cert-manager.yaml")
  ]
}

/*variable "provider-eks" {
  default = "kubernetes.eks"
}*/

# Nginx Ingress Controller

// https://registry.terraform.io/modules/byuoitav/nginx-ingress-controller/kubernetes/latest
module "nginx-ingress-controller" {
  source  = "byuoitav/nginx-ingress-controller/kubernetes"
  version = "0.2.1"
  # insert the 1 required variable here
}


/*module "nginx-controller" {
  # source  = "terraform-iaac/nginx-controller/helm"
  source  = "terraform-iaac/nginx-controller/kubernetes"
  version = "1.3.1"

  additional_set = [
    {
      name  = "controller.service.annotations.service\\.beta\\.kubernetes\\.io/aws-load-balancer-type"
      value = "nlb"
      type  = "string"
    },
    {
      name  = "controller.service.annotations.service\\.beta\\.kubernetes\\.io/aws-load-balancer-cross-zone-load-balancing-enabled"
      value = "true"
      type  = "string"
    }
  ]
}*/

# ALB Ingress Controller
/*module "alb_ingress_controller" {
  source  = "iplabs/alb-ingress-controller/kubernetes"
  version = "3.1.0"

  providers = {
    kubernetes = var.provider-eks
  }

  k8s_cluster_type = "eks"
  k8s_namespace    = "kube-system"

  aws_region_name  = var.aws_region
  k8s_cluster_name = data.aws_eks_cluster.cluster.name
}*/

# TODO: remove external-dns from the list and instead follow: https://github.com/hashicorp/terraform-provider-kubernetes-alpha/issues/72
resource "helm_release" "helm-releases" {
  count = length(var.services)
  name = var.services[count.index]

  # Ensure service account exists before attempted deployment of external-dns
  depends_on = [
    kubernetes_service_account.external-dns,
    kubernetes_namespace.cert-manager,
    helm_release.cert-manager,

    kubernetes_cluster_role.external-dns,
    kubernetes_cluster_role_binding.external-dns-viewer,
    aws_iam_role.external-dns
  ]

  values = [
    file("../env/charts/values-prod.yaml")
  ]

  chart = "${var.relativeChartPath}/${var.services[count.index]}"
}

# Deploy external-dns

# TODO: Create cert-manager namespace
/*resource "kubernetes_manifest" "cert-manager-cluster-issuer" {
  provider = kubernetes
  manifest = yamldecode(templatefile("manifests/cert-manager-cluster-issuer.tmpl.yml", {
    cluster_name=var.cluster_name
  }))
  depends_on = [helm_release.cert-manager]
}*/
