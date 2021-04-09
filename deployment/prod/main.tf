/*
provider "aws" {
  region = "us-west-2"
  shared_credentials_file = "/Users/stevenelleman/.aws/credentials"
}
*/

# Same config parameters as k8s terraform parameter
provider "kubectl" {
  #host                   = var.eks_cluster_endpoint
  #cluster_ca_certificate = base64decode(var.eks_cluster_ca)
  #token                  = data.aws_eks_cluster_auth.main.token
  #load_config_file       = false
}

variable "services" {
  description = "Service names"
  default = ["api-store", "grpc", "public-api", "web-frontend", "external-dns"]
}

# Loop over services to load per-service yaml contents
data "kubectl_file_documents" "manifest" {
  count = length(var.services)
  content = file("../../services/${var.services[count.index]}/k8s.yaml")
}

# Merge service documents
locals {
  resources = flatten([
  for svc in data.kubectl_file_documents.manifest  : [
    svc.documents
  ]
  ])
}

# Deploy service resources
resource "kubectl_manifest" "resource-deployments" {
  count = length(local.resources)
  yaml_body = element(local.resources, count.index)
}
