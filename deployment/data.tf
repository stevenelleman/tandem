data "aws_eks_cluster" "cluster" {
  name = module.eks.cluster_id
}

data "aws_eks_cluster_auth" "cluster" {
  name = module.eks.cluster_id
}

data "tls_certificate" "cluster" {
  url = data.aws_eks_cluster.cluster.identity.0.oidc.0.issuer
}

// Define Trust Relationship between OIDC Provider and Role
# Source: https://docs.aws.amazon.com/eks/latest/userguide/iam-roles-for-service-accounts-technical-overview.html
data "template_file" "role_trust_relationship" {
  template = file("${path.module}/policies/role_trust_relationship.json.tmpl")
  vars = {
    oidc_arn      = aws_iam_openid_connect_provider.cluster.arn
    oidc_url      = aws_iam_openid_connect_provider.cluster.url
    k8s_namespace = var.k8s_namespace
    service_account_name     = var.external_dns_service_account_name
  }
}