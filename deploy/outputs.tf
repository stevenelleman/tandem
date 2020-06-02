output "private_ips" {
  description = "NGINX Private IPs"
  value       = aws_instance.web.private_ip
}

output "private_dns" {
  description = "NGINX Private DNS"
  value       = aws_instance.web.private_dns
}

output "public_ips" {
  description = "NGINX Public IPs"
  value       = aws_instance.web.public_ip
}

output "public_dns" {
  description = "NGINX Public DNS"
  value       = aws_instance.web.public_dns
}