# sg

## Docs
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

## Tasks 
- [x] Initial Deployment Process [Tuesday]
- [x] Frontend w/ Mock Endpoint/Data [Tuesday]
- [x] API Service w/ Mock Data [Wednesday]
- [x] Postgres DB [Thursday]
- [x] Frontend Linter [Thursday]
- [x] Redux State Management [Friday]
- [x] Web-Frontend Client [Friday]
- [x] Connect DB to the Frontend [Friday]
- [x] Frontend Jest Testing [Sunday]
- [x] React-Router [Sunday]
- [ ] API Table Migrations [Sunday]
- [ ] Local VM Setup (Using Tilt preferably) [Sunday]
- [ ] Backend Go Linter
- [ ] API Service Test Pattern
- [ ] Server-Side Transformations

Backlog:
- [ ] Build Process for Dashboard and API service
- [ ] Microservice Deployment Process
- [ ] Authz Context and Login
    - [ ] Browser Cookier Manager 
    - [ ] User Info Endpoint
- [ ] Frontend Redux Testing
- [ ] API Service Error Handling 
- [ ] API Service Logging
- [ ] Per-Request DB Connection Instantiation
