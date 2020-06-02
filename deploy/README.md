# Deploy 

## Setup 
1. (For intellij) Download the Hashicorp Terraform Plugin and restart the IDE.
2. For macOS, install `terraform` CLI with `brew install terraform`. For other OSes download and install [`terraform` CLI](https://www.terraform.io/downloads.html). This tool will be used for executing terraform locally.
3. Download and install [`aws-cli`](https://docs.aws.amazon.com/cli/latest/userguide/install-cliv2.html). This tool will be used for executing terraform files on an associated AWS account.
4. In [AWS security credential console](https://console.aws.amazon.com/iam/home#/security_credentials) navigate to access keys and generate a new access key.
5. Locally create a `credentials` file, the default location is `$HOME/.aws/credentials`.
 ```
sudo vim ~/.aws/credentials
```
The format of the file should be as follows: 
```
[${profilename, set to default}]
aws_access_key_id=${key id from console}
aws_secret_access_key=${secret key from console}
```
Where `${...}` indicates a variable. Other info can be found [here](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-files.html)
6. Run `terraform plan`. If everything has been set up properly, the call should be successful.
7. Run `terraform apply` to spin up updated web service.


## Tips
- `TF_LOG=debug terraform plan` is your best friend.