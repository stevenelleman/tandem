# External DNS 

## What is External DNS? 
External DNS manages your provider's Hosted Zone. A Hosted Zone maps a domain to name servers and serves DNS records. External DNS modifies these DNS to be pointing to specific k8s Service IP. 

In our case we set `--domain-filter=< domain >` and `--txt-owner-id=< hosted zone id >` so that external dns manages the specified hosted zone and associates it to services with matching domain. In the web-frontend service  we set `external-dns.alpha.kubernetes.io/hostname` to `www.grouphouse.io` for its target domain and the filter picks up on this value so that the records are mapped to the web-frontend service. 

To enable External DNS to modify records we will need to associate an IAM policy via a user or role (I used a user for the time being) so that it can act on behalf of the user or assume the role. 

## Setup 
1. Define AWS IAM policy to enable external DNS to update the hosted zone records. Replace the hosted zone id. 
```
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": "route53:ChangeResourceRecordSets",
            "Resource": "arn:aws:route53:::hostedzone/< hosted zone id >"
        },
        {
            "Effect": "Allow",
            "Action": [
                "route53:ListHostedZones",
                "route53:ListResourceRecordSets"
            ],
            "Resource": "*"
        }
    ]
}
```
2. Make AWS IAM user using this policy. 
3. Locally run: `kubectl create secret generic test-secret --from-literal=AWS_ACCESS_KEY_ID='< access key id >' --from-literal=AWS_SECRET_ACCESS_KEY='< access key >' --from-literal=AWS_HOSTED_ZONE_ID='< your target hosted zone id >'` using the credential info of this user. `envFrom.secretRef` will pick up on this key and run as the user.  
4. Update the `--domain-filter` and `external-dns.alpha.kubernetes.io/hostname` to whatever the hostname of the hosted zone is. 