# sg

## Docs
- TODO ([./TODO.md](./TODO.md))
- Spec ([./docs/spec.md](./docs/spec.md))
- App Setup Docs ([./services/web-frontend/README.md](./services/web-frontend/README.md))
- Service Deployment ([./deploy/README.md](./deploy/README.md))

Services: 
- Web Frontend Service ([./services/web-frontend/src/components/README.md](./services/web-frontend/src/components/README.md))
- Public API Service ([./services/public-api/README.md](./services/public-api/README.md))

## Docker Setup
Run `docker-compose up`

## Code-Amendment Best Practices 
- Don't commit directly to master, make PRs. 
- PRs should be reviewed and be required to have a +1 from non-contributor to the PR.
- PRs should have corresponding tests. 
- PRs should include documentation. Documentation should be kept with its relevant code to minimize drift.
- Correct known trip-wires (little annoyances, mistakes, gotchas) immediately - quick fixes pay off dividends.
- Prefix branch name with creator's initials.