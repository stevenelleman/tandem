# sg

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
Run `tilt up`
 
Or run `docker-compose up`

## Services: 
- [Web Frontend Service](./services/web-frontend/src/components/README.md)
- [Public API Service](./services/public-api/README.md)
- [Postgres API Store](./services/api-store/README.md)
- [Sourcegraph Service](./services/api-store/README.md)

## Docs
- [TODO](./TODO.md)
- [Spec](./docs/spec.md)
- [Libraries](./libraries/golang/README.md)
- [Service Deployment](./deploy/README.md)
