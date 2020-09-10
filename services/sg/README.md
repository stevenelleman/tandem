# Sourcegraph Service 

## Table of Contents 
- [Objective](#objective)
- [Questions](#questions)
- [Requirements](#requirements)
- [Service Interface](#interface)
- [Database](#database)

## Docker 
1. Run `docker build -t sg .` in `./sg` 

## Objective
Offer a separate service (and database) for transforming sources into scopes. The service will not be publicly accessible. It will be accessed through the `Public-API`.

## Questions
- What DB type should be used? 
- Any DB for functions? How will the functional transformations be saved? 

## Relationships 
- `Public-API` service will public interface to service -- will grpc client beneath the hood for improved performance. 

## Requirements
1. Text: Only supports text file types.  
2. Local: Only constructs sourcegraph on a single computer. 
3. Bidirectional: Links can go both ways. 
4. Persistent: Links cannot disappear. 
5. Pure Transformations: Transformations should be pure functions to ensure documents are deterministic. 
6. Invertible Transformations: Only if the transformation is inverted can it be unapplied efficiently. 
6. Version-controlled: By default the history of changes is saved.
7. Scoped: Sources can be organized into collections. A document is one type of scope, it is a collection of connected sources. 
8. Set Logic: Set logic can be applied to collections, such that new scopes can be derived from existing scopes.  

## Interface
Use GRPC interface. 

## Database 
GraphQL? 
