# TODO 

## Table of Contents 
- [Tasks](#todo)
- [Backlog](#backlog)
- [Known Unknowns](#known-unknowns)    
- [Answered](#answered)   
- [Completed](#completed)

## Tasks 
From highest priority to lowest.
UI: 
- [ ] Mapping Actions to Transformations 
- [ ] Mapping UI to Actions 
- [ ] Managing requests from multiple Panes

Architecture: 
- [ ] Doc chunking design 
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

Public-API: 
- [ ] Health endpoint (https://blog.gopheracademy.com/advent-2017/kubernetes-ready-service/)
- [ ] Graceful shutdown (https://blog.gopheracademy.com/advent-2017/kubernetes-ready-service/)
- [ ] Request JSON-to-Struct Mapping
- [ ] Backend Go Linter
- [ ] Pagination pattern
- [ ] Jaegar spanner
- [ ] API Service Test Pattern
- [ ] API Service Error Handling 
- [ ] API Service Logging

Other: 
- [ ] Clean up READMEs 

## Backlog
- [ ] Migrate command
- [ ] Build Process for Dashboard and API service
- [ ] Microservice Deployment Process
- [ ] Authz Context and Login
    - [ ] Browser Cookier Manager 
    - [ ] User Info Endpoint
- [ ] Frontend Redux Testing
- [ ] Per-Request DB Connection Instantiation
- [ ] Feature-flagging -- what's the proper level? 

## Future Backlog 
- [ ] Redis: Map session id to cached chunks 


## Known Unknowns
- How can scopes express collections of other scopes? Or should that involve a separate abstraction? 
- Data preprocessing is extremely expensive for ML models. This is a known opportunity. How can sourcegraph be organized to act as a foundation for ML models? 
- Is it possible to produce end-documents from scopes describing scopes? 
- Parent-child scopes? Ex: document scope vs chapter scope. Should the chapter scope exist within the document? How should scopes be linked? 
- Should the UI interface also be a sourcegraph? 
- What dictates whether two scopes can interact? An overlap in transformations?
- Should go mod be initialized in each golang service, as opposed in the root ./sg directory?

## Answered
- What is an efficient way to apply all the sources to find the current state of a doc?
    - Separate service. Separate service is responsible for 
- How to offer different UI interfaces? 
    - This is where Mike's approach may be easier, having allowing other websites. 
    - Middleground is allowing plugins. 
- Where do metrics fit into the social abstractions? 
    - Metrics are for credibility. Credibility is tied to trust. Trust is tied to silos.
- Examples of metrics: 
    - Minimum and maximum metrics: memberships and awards.
- Should scopes have to be owned by a silo? 
    - Yes. 
- Can scopes be able to change owners? 
    - No. Creator-status is involiable. 
- Scope vs overlay? 
    Scope is chain of transformations. Overlay are "surrounding" documents that reference _it_.     
- Should faceted identities be allowed to own multiple silos?
    - Each faceted identity should have its own personal silo. 
    - I see no downside in allowing multiple personal silos for a faceted identity. 
    - I see no downside in allowing a single silow for all a person's faceted identity. 
- Should multiple silos be able to own one scope? 
    - No. Fork the scope for the other silo. 
- When should a document be saved? Should it be explicit based on pressing publish button or should it be implicit via auto-saves? 
    - Explicit. I like picturing sourcegraph as an IDE. Whenever command-S is pressed the current diff is saved as a source.
- Should it be possible to decouple an edge and source? Is there any use-value to that? 
    - An edge is a transformation. To produce the end-document it's necessary. Edges are necessary at the scope-level. Sources can be self-contained and still be meaningful
- Should it be possible to delete history?
    - No, only soft deletions are allowed.  
- Should a starting scope function take an empty argument? 
    - Yes, initial doc should be null.  
- Should one silo be allowed to own another silo? Ex: a company vs team. 
    - It should be a feature but by default an organization should be like a microservice - interdepedent but decoupled. 
- Default silos? 
    - History of read/writes

## Completed
From most recent to oldest.
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
