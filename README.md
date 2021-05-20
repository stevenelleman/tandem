# Web Microservice Shell

## Code-Amendment Best Practices 
- Don't commit directly to master, make PRs. 
- PRs should be reviewed and be required to have a +1 from non-contributor to the PR.
- PRs should have corresponding tests. 
- PRs should include documentation. Documentation should be kept with its relevant code to minimize drift.
- Correct known trip-wires (little annoyances, mistakes, gotchas) immediately - quick fixes pay off dividends.
- Prefix branch name with creator's initials.
- As new items come up, add to [TODO](./TODO.md). 
- Maximize code reuse, minimize service directory size and complexity

## Development Setup
Run `tilt up` 

and navigate to `localhost` (no port).

You can also run the services via `docker-compose up`, individually via dockerfile, or in terminal.

## TF Deployment on AWS
[Deployment Docs](./deployment/prod)

1. Navigate to `./deployment/prod`
2. Run `terraform init` to download providers.
3. `terraform plan`
4. `terraform apply`
5. Update kubectl config: `aws eks --region=us-west-1 update-kubeconfig --name=${cluster name}`

## Charts:
- [External DNS README](charts/external-dns/README.md)
- [Web Frontend Service](charts/web-frontend/Chart.yaml)
- [Public API Service](charts/public-api/Chart.yaml)
- [Example GRPC Service](charts/grpc/Chart.yaml)
- [Postgres API Store](charts/api-store/Chart.yaml)

## Dockerized Services:
- [Web Frontend Service](./services/web-frontend/src/components/README.md)
- [Public API Service](./services/public-api/README.md)
- [Example GRPC Service](./services/grpc)

## Docs
- [TODO](./TODO.md)
- [Libraries](./libraries/golang/README.md)
- [Service Deployment](deployment/website/README.md)
