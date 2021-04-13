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

## Setup
Change the domain to your target domain in these files: 
- [Domain Filter in external-dns service](./services/external-dns/k8s.yaml)
- [Hostname annotation in web-frontend service](./services/web-frontend/k8s.yaml)
- [REACT_APP_HOST in web-frontend dockerfile](./services/web-frontend/Dockerfile)
- [TiltHost in public-api service](./services/public-api/constants/constants.go)

Follow [`external-dns` instructions](./services/external-dns/README.md) if you intend to use it.

Run `tilt up` 

and navigate to your target domain (no port). 

You can also run the services via `docker-compose up`, individually via dockerfile, or in terminal.

## Services: 
- [External DNS](./services/external-dns/README.md)
- [Web Frontend Service](./services/web-frontend/src/components/README.md)
- [Public API Service](./services/public-api/README.md)
- [Postgres API Store](./services/api-store/README.md)
- [Example GRPC Service](./services/grpc)

## Docs
- [TODO](./TODO.md)
- [Libraries](./libraries/golang/README.md)
- [Service Deployment](deployment/website/README.md)
