# Public API Service 

## Table of Contents 
- [Local Setup](#code-setup)
- [Docker Setup](#docker-setup)
    - [Postgres Setup](#postgres)
    - [API Setup](#api)
- [Questions](#questions)
- [Layers](#layers)
- [API Routes](#api-routes)

## Code Setup
1. Intellij File-Watcher: gofmt and goimports
1. Run `go mod tidy`. This will download all necessary go packages. 
2. Run `go build` in `api` directory. 
3. Run `./api` in directory. 
4. ~ Postgres setup ~ 
5. Access `psql` db via `psql stevenelleman`

## Docker Setup

### API  
1. Run `docker build -t sg_public_api .` in `./public-api` 
2. Run `docker run -p 8000:8000 public-api`

Useful Curl Commands:
- `curl -X POST localhost:8000/v1/silos/${id} -H 'Content-Type:application/json' -d '{"state":"active"}'`: To make a REST request to the API service.

## Questions
- Capitalized properties or now in models? 
- Should scopes always be associated with a silo?
    - I generally like the idea of a personal silo, but this begs the question: can a scope have multiple parent silos? 
    - I think this is something you have to play around with.
- Type mappings are required at two interfaces: the API and the DBI. Some or many of these mappings will be used for both mappings, but it would be nice to have a logical separation. For the DB `models` seems to communicate what's going on. 

## TODO: 
- [ ] Table migration into PSQL 
- [ ] API testing 
- [ ] Map request object

## Layers 

[API](./api) -request-> [Handler](../../libraries/golang/guts/handlers) -req-> [Controller](../../libraries/golang/guts/controller) -req-> [DBI](../../libraries/golang/guts/connection) -req-> [DB](../../libraries/golang/guts/connection/service/psql_conn/db) 

API <-resp- Handlers <-resp- Controllers <-resp- DBI <-resp- DB

1. `API`: Public interface to application. The interface should be seen as an inviolable contract with the public and our users. The _whole point_ of an interface is to be fixed so that the _complexity behind it_ can be flexible.  
2. `Handler`: Methods corresponding to `API` endpoints. `Handlers` _use_ `Controller` methods to interact with the `DB`. The `Handler` will also handle logging, as it's the "top-level" 
3. `Controller`: `Controller` methods use `DBI` methods to interact with the `DB` Additional layer of indirection for greater flexibility and reuseability. 
4. `Database Interface (DBI)`: Methods that act as an interface with the `DB`. Allows the ability to input queries to the `DB` without using the query language itself. 
5. `Database (DB)`: Postgres database. 