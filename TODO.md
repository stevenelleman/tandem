# TODO 

## Table of Contents 
- [Tasks](#todo)
- [Backlog](#backlog)
- [Known Unknowns](#known-unknowns)      
- [Completed](#completed)

Unordered Tasks: 
- [ ] Fix go mod pathing

Architecture: 
- [ ] Connection/ConnectionFactory interface 
- [ ] Set up graph db source-store 
    - [ ] Think of pattern for sg service and source-store
        - [ ] Interface 
        - [ ] Layering 
        - [ ] Use-Cases / Requirements 
- [ ] Where to apply transformations? How to represent? 
    - Would like some kind of function-db where functions are dumbly applied to objects
- [ ] Representing Transformations, Actions, and UI as graph-doc 

Environment:
- [ ] Tilt: Convert to using DNS names -- finally connected the web-frontend 
        and public-api by opening public-api -- need to make it private 
- [ ] Envoy (https://www.envoyproxy.io/)

Public-API: 
- [ ] Health endpoint (https://blog.gopheracademy.com/advent-2017/kubernetes-ready-service/)
- [ ] Graceful shutdown (https://blog.gopheracademy.com/advent-2017/kubernetes-ready-service/)
- [ ] Request JSON-to-Struct Mapping
- [ ] Authz 
    - [ ] Require token with all requests + API calls 
    - [ ] Session object 
- [ ] Backend Go Linter
- [ ] Pagination pattern
- [ ] Jaegar spanner
- [ ] API Service Test Pattern
- [ ] API Service Error Handling 
- [ ] API Service Logging

K8s/Deployment: 
- [ ] Testing Public Zone for frontend 
- [ ] Switch to port 443 for frontend 
- [ ] Per-commit automated deployments 

Other: 
- [ ] Clean up READMEs 

## Backlog
- [ ] Migrate command
- [ ] Per-Request DB Connection Instantiation
- [ ] Feature-flagging -- what's the proper level? 

## Future Backlog 
- [ ] Redis: Map session id to cached chunks 

## Known Unknowns


## Completed
From most recent to oldest.
- [x] Cleaner client connection creation 
- [x] Move Handler-Controller-Connection to libraries 
    need to pass in store information at beginning, if provided then connection is initialized. 
    alternatively could move out a copy and play there 
    need to imagine what this would look like and how complex it is. Upside is that there would a common way of init-ing golang services 
    downside is "one-size fits all". Requires better layered interface before moving it out into an external library.
- [x] Running kubernetes-orchestrated cluster
- [x] Wire ctx through public-api
- [x] Tilt Setup: 
    - [x] Convert docker-compose over
    - [x] Be able to access web-frontend 
    - [x] Public-api run 
    - [x] web-frontend can hit public-api 
    - [x] Public-api initialized api-store
    - [x] Add sg service
- [x] Set up grpc client for sourcegraph service
- [x] Use squirrel querybuilder 
- [x] API Layering abstraction
- [x] Shared Code and Vendored Library Pattern
- [x] Mapper
- [x] API Table Migrations
- [x] React-Router
- [x] Frontend Jest Testing
- [x] Connect DB to the Frontend
- [x] Web-Frontend Client
- [x] Redux State Management
- [x] Frontend Linter
- [x] Postgres DB
- [x] API Service w/ Mock Data
- [x] Frontend w/ Mock Endpoint/Data
- [x] Initial Deployment Process
