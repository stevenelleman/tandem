resource "aws_iam_openid_connect_provider" "cluster" {
  url = data.aws_eks_cluster.cluster.identity.0.oidc.0.issuer

  client_id_list = [
    "sts.amazonaws.com",
  ]

  thumbprint_list = [
    data.tls_certificate.cluster.certificates.0.sha1_fingerprint
  ]
}

// External DNS Policy
resource "aws_iam_policy" "external-dns-policy" {
  name        = var.external_dns_policy_name
  description = "Policy that allows iam service account to updated DNS records of public hosted zone"
  path        = "/"
  policy      = file("${path.module}/policies/allow_external_dns_updates_policy.json")
}

// Create External-DNS Role
resource "aws_iam_role" "external-dns" {
  // TODO: change back to role name
  name   = var.external_dns_service_account_name
  path   = "/"

  # Source: https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_role
  assume_role_policy = data.template_file.role_trust_relationship.rendered

  # Source: https://dzone.com/articles/how-to-use-aws-iam-role-on-aws-eks-pods
  force_detach_policies = false
}

// Attach External DNS Policy to Role
resource "aws_iam_policy_attachment" "attach-external-dns-update" {
  name = "attach-external-dns-update"
  roles      = [aws_iam_role.external-dns.name]
  policy_arn = aws_iam_policy.external-dns-policy.arn
}

// Prod Cluster Role
resource "kubernetes_service_account" "external-dns" {
  metadata {
    name = var.external_dns_service_account_name
    namespace = var.k8s_namespace
    annotations = {
      // Assign External-DNS Role to Service Account
      "eks.amazonaws.com/role-arn": aws_iam_role.external-dns.arn
    }
  }
  automount_service_account_token = true
}

variable cluster_role_name {
  default = "external-dns"
}

// K8s Cluster Role and Binding for Service Account
resource "kubernetes_cluster_role" "external-dns" {
  metadata {
    name = var.cluster_role_name
  }

  rule {
    api_groups = [""]
    resources = ["services","endpoints","pods"]
    verbs      = ["get", "list", "watch"]
  }

  rule {
    api_groups = ["extensions","networking.k8s.io"]
    resources = ["ingresses"]
    verbs      = ["get", "list", "watch"]
  }

  rule {
    api_groups = [""]
    resources = ["nodes"]
    verbs      = ["list","watch"]
  }
}

resource "kubernetes_cluster_role_binding" "external-dns-viewer" {
  metadata {
    name = "external-dns-viewer"
  }
  role_ref {
    api_group = "rbac.authorization.k8s.io"
    kind      = "ClusterRole"
    name      = var.cluster_role_name
  }
  subject {
    kind      = "ServiceAccount"
    name      = var.external_dns_service_account_name
    namespace = var.k8s_namespace
  }
}
