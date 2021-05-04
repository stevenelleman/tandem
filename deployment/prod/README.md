# AWS Terraform Deployment 

# Setup: 
1. Change local `project_name` variable in `var.tf` to your project. 
2. Run `terraform init` to download providers. 

## Overview 

Currently web-frontend and external-dns work. This is the setup for how to deploy it. 

https://www.freecodecamp.org/news/how-to-setup-dns-for-a-website-using-kubernetes-eks-and-nginx/ was immensely helpful for figuring out the issue. 
https://dzone.com/articles/how-to-use-aws-iam-role-on-aws-eks-pods

## New Deployment 

```
terraform plan 

terraform apply

aws eks --region=us-west-1 update-kubeconfig --name=${cluster name}
```

## Old Deployment
1. Run `terraform apply` w/o `external-dns` service (i.e. remove from service list in `main.tf`). 
Note the cluster name. 

2. Updated kubectl for new cluster: `eksctl get cluster --region=us-west-1` and `aws eks --region=us-west-1 update-kubeconfig --name=<new cluster>`.

2. After completion, make iam policy: 
```
aws iam create-policy \
  --policy-name AllowExternalDNSUpdates \
  --policy-document '{"Version":"2012-10-17","Statement":[{"Effect":"Allow","Action":["route53:ChangeResourceRecordSets"],"Resource":["arn:aws:route53:::hostedzone/*"]},{"Effect":"Allow","Action":["route53:ListHostedZones","route53:ListResourceRecordSets"],"Resource":["*"]}]}'
```
Note the created policy arn. 

3. Make iam service account associated with the new iam policy: 
```
eksctl create iamserviceaccount --cluster=your-cluster-name --name=external-dns --namespace=default --attach-policy-arn=arn:aws:iam::123456789:policy/AllowExternalDNSUpdates
```
This command will generate a new iam role with a name like: `arn:aws:iam::123456789:role/eksctl-grouphouse-eks-6oHtMVpZ-addon-iamserv-Role1-1R0SEFHJ8Z7PN`. Note it for later. 

4. Attach generated iam role name to external-dns service account `eks.amazonaws.com/role-arn`. 

5. Uncomment out `external-dns` service in `main.tf` and deploy the associated resources. 

Done! 

## What's Next 

The crux of the issue is getting the generated role-arn to associate with the `external-dns` service account. This generated role has special selection conditions for when to apply the new iam policy, such as: 
```system:serviceaccount:default:external-dns``` is designated a trusted entity. 
