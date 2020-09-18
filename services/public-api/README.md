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


1. `API`: Public interface to Sourcegraph. The interface should be seen as an inviolable contract with the public and our users. The _whole point_ of an interface is to be fixed so that the _complexity behind it_ can be flexible.  
2. `Handler`: Methods corresponding to `API` endpoints. `Handlers` _use_ `Controller` methods to interact with the `DB`. The `Handler` will also handle logging, as it's the "top-level" 
3. `Controller`: `Controller` methods use `DBI` methods to interact with the `DB` Additional layer of indirection for greater flexibility and reuseability. 
4. `Database Interface (DBI)`: Methods that act as an interface with the `DB`. Allows the ability to input queries to the `DB` without using the query language itself. 
5. `Database (DB)`: Postgres database. 

## API Routes 

Silos: An authorization context must be included that includes a Faceted Identity (FI) of a silo member. All child routes must also include corresponding FI. 
- `GET`-`v1/silos`: List silos of FIs.  
- `GET`-`v1/silos/${id}`: Get silo details.
- `POST`-`v1/silos/${id}`: Create silo.  
- `PUT`-`v1/silos/${id}`: Edit silo. 
- `DELETE`-`v1/silos/${id}`: Soft-delete silo. Need to decide what happens to scopes - are they still accessible by members? Are they publicized? 

Silo FIs: 
- `GET`-`v1/silos/${id}/scopes/${id}/members`: List of members of silo. 
- `GET`-`v1/silos/${id}/scopes/${id}/members/${id}`: Get member details of silo. 
- `POST`-`v1/silos/${id}/scopes/${id}/members/${id}`: Add member.
- `PUT`-`v1/silos/${id}/scopes/${id}/members/${id}`: Edit member.
- `DELETE`-`v1/silos/${id}/scopes/${id}/members/${id}`: Remove member.

Scopes: 
- `GET`-`v1/silos/${id}/scopes`: List of scopes corresponding to silo. 
- `GET`-`v1/silos/${id}/scopes/${id}`: Get scope details.
- `POST`-`v1/silos/${id}/scopes/${id}`: Create scope.
- `PUT`-`v1/silos/${id}/scopes/${id}`: Edit scope details.
- `DELETE`-`v1/silos/${id}/scopes/${id}`: Soft-delete scope.

Scope Versions: 
- `GET`-`v1/silos/${id}/scopes/${id}/version/${int}`
- `POST`-`v1/silos/${id}/scopes/${id}/version/${int}`: Create new scope version. Version must be n + 1, where n is the current last version. 
No edit or delete endpoints. Versions are immutable. 

Scope FIs: 
- `GET`-`v1/silos/${id}/scopes/${id}/members`: List contributors. 
- `GET`-`v1/silos/${id}/scopes/${id}/members/${id}`: List contributor details. 
No edit or delete endpoints.

Forums: 
- `GET`-`v1/forums`: List all forums corresponding to FIs.
- `GET`-`v1/forums/${id}`: Get forums details.
- `POST`-`v1/forums/${id}`: Create forum.
- `PUT`-`v1/forums/${id}`: Update forums details.
- `DELETE`-`v1/forums/${id}`: Soft delete forum.

Forum silos: 
- `GET`-`v1/forums/${id}/silos`: List all silos belonging to forum.
- `GET`-`v1/forums/${id}/silos/${id}`: Get silo details belonging to forum.
- `POST`-`v1/forums/${id}/silos/${id}`: Add silo to forum.
- `PUT`-`v1/forums/${id}/silos/${id}`: Update silo details.
- `DELETE`-`v1/forums/${id}/silos/${id}`: Remove silo from froum. 

Forum FIs: 
- `GET`-`v1/forums/${id}/members`: List all members belonging to forum.
- `GET`-`v1/forums/${id}/members/${id}`: Get members details.
- `PUT`-`v1/forums/${id}/members/${id}`: Update silo details, in particular the ability to soft-exile/delete a member from a forum. 
No create or delete endpoints. Create or delete are controlled at the silo-level.  