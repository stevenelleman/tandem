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

## Versions 
### V1: MVP 
Maximally dumb. 
1. Pull the entire document initially 
2. Change historical versions pulls the entire document (which would likely work on small documents)

## Use-Cases 
- Reader is advancing through doc and needs to pull more chunks 
    [Size of Pane determines how much content can fit -- need some kind of sizing interface that can be applied for different content types -- put first on text. For text just needs to be frontsize, height, width and starting character index -- can calculate the end based on this]
- Reader is going through history of doc and needs to get chunk for correct version -- suggests some kind of caching for that current version, don't want to have to recalculate the doc every time 
    [Alternatively can naively store chunk every version]

## Implementation Ideas 
1. Chunking Cases: 
- Same version:  
    - Have a version of the doc, only send some window of the doc. 
- Between different versions: 
    - How to keep same position, as perceived by person? Assume that focus is at the top of the page, stick with that specific text, until it's replaced or deleted. When replaced, switch that using text that spatially replaced it. When deleted, use text immediately before or after to fix  
- Translate doc from one version to the next? 
    - Need some kind of patching method -- very similar to the transformation function itself -- maybe the same thing? Would be convenient 
    - Alternatively could cache remotely for some kind of session id -- that makes the most sense to me 
    - Would need some kind of caching service that's really a key-value store between a session id and the current version (Redis?) 
    - Could put the client in the controller 
    - No need for this initially, initially just prep some kind of window object and redo all the logic until it's prohibitive -- with text it would probably would be too prohibitive provided some kind of chunking is used
    - How to decouple session id with FIs? 
        - But session id into each FI cookie, session id just maps to the chunks
    - Any way to "restore" what was saved from previous session? Any security concerns? 
    - session 
    
2. Transformations 
- How to represent the inverse? 
    - fxn(doc0, src0) => doc1
    - fxn^-1(doc1) => doc0, src0, probably the easiest way is simply saving the version and pulling, when going back may not have src0 
    
3. Resolving Commits: [Necessary for social version] 

## Interface 

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
GraphQL? -- Seems like most appropriate db type.  
