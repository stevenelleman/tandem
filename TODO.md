# TODO 

## Overview 
These are TODO items for the generic Tandem library 

## Table of Contents 
- [Tasks](#todo)
- [Backlog](#backlog)
- [Known Unknowns](#known-unknowns)      
- [Completed](#completed)

Architecture: 
- [ ] Where to apply transformations? How to represent? 
    - Would like some kind of function-db where functions are dumbly applied to objects
- [ ] Representing Transformations, Actions, and UI as graph-doc 

K8s/Deployment: 
- [ ] Investigate Kustomize and migrate if preferred to Helm 
- [ ] Envoy (https://www.envoyproxy.io/)
- [ ] Per-commit automated deployments 

Web-Frontend: 
- [ ] Move client out to typescript library 

Public-API: 
- [ ] Retry establishing API-Store connection, or delay slightly 
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

API-Store: 
- [ ] Simply yaml setup, probably should make as independent from k8s as possible. 
- [ ] Figure out why persistent volume and persistent volume claim are sometimes not created on `tilt up`.

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
- [x] Connection/ConnectionFactory interface [Opted for: shared controller has multiple handlers, controller type is specific to data store combo]
- [x] Tilt: Convert to using DNS names -- finally connected the web-frontend 
        and public-api by opening public-api -- need to make it private 
- [x] Testing Public Zone for frontend 
- [x] Switch to port 443 for frontend 
- [x] Fix go mod pathing
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
