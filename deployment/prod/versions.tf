terraform {
  required_version = ">= 0.13"

  required_providers {
    kubectl = {
      source = "gavinbunney/kubectl"
      version = "1.10.0"
    }

    aws = {
      source  = "hashicorp/aws"
      version = "3.35.0"
    }

    tls = {
      source = "hashicorp/tls"
      version = "3.1.0"
    }

    local = {
      source  = "hashicorp/local"
      version = "2.1.0"
    }

    kubernetes = {
      source  = "hashicorp/kubernetes"
      version = "2.0.3"
    }

    null = {
      source  = "hashicorp/null"
      version = "3.1.0"
    }

    random = {
      source  = "hashicorp/random"
      version = "3.1.0"
    }

    template = {
      source  = "hashicorp/template"
      version = "2.2.0"
    }
  }
}