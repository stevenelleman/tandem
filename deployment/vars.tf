variable "aws_region" {
  default = "us-west-1"
}

variable "project_name" {
  default = "grouphouse"
}

variable "aws_account_id" {
  default = "016489005889"
}

variable "credential_source" {
  default = "/Users/stevenelleman/.aws/credentials"
}

variable "k8s_namespace" {
  default = "default"
}

variable "external_dns_service_account_name" {
  # Must match `serviceAccountName` value in `services/external-dns/k8s.yaml`
  default = "external-dns"
}

variable "external_dns_role_name" {
  default = "external-dns-iam-role001"
}

variable "external_dns_policy_name" {
  default = "external-dns-iam-policy001"
}