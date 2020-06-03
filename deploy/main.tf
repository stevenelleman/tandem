variable image_id {
  type = string
  default = "ami-039a49e70ea773ffc"
  description = "AMI id, Ubuntu Xenial 16.04 by default"
}

variable instance_type {
  type = string
  default = "t2.micro"
}

// security group
variable sourcegraph_sg {
  type = string
  default = "sg-03b567b3c3a72f196"
  description = "Security group to use, default group allows HTTPS (open), HTTP (open), and SSH (home ip)"
}

variable subnet_id {
  type = string
  default = "subnet-7d5ec043"
  description = "Subnet in VPC (virtual private cloud) to use. VPC is default."
}

variable key_name {
  type = string
  default = "sourcegraph-key"
  description = "Key pair name that will be used with created EC2 instance."
}

provider "aws" {
  region = "us-east-1"
  shared_credentials_file = "$HOME/.aws/credentials" // default credentials value, just wanted to be explicit
}

resource "aws_instance" "web" {
  ami = var.image_id
  instance_type = var.instance_type
  key_name = var.key_name
  vpc_security_group_ids = [var.sourcegraph_sg]
  subnet_id = var.subnet_id

  user_data = templatefile("./userdata.tmpl", {})
}

resource "aws_eip_association" "eip_assoc" {
  instance_id   = aws_instance.web.id
  allocation_id = "eipalloc-0cabdab2d8aec5451" // Set to existing aws_eip to keep the public_ip the same
}
