# Frontend

## Table of Contents 
- [Routes](#frontend-routes)
- [Views](#views)
- [Directory Structure](#directory-structure)
- [Chunking](#chunking)
- [Separation of Concerns](#separation-of-concerns)
- [Faceted Identities](#faceted-identities)
- [Routing](#routing)
- [Timeline](#timeline)
    - [Timeline Modes](#modes)
- [Open Source Interfaces](#interface-for-interfaces)
- [Login](#login)

## Frontend Routes  
- `www.sourcegraph.wtf/v1/documents`
- `www.sourcegraph.wtf/v1/scopes` -> show all available scopes   
- `www.sourcegraph.wtf/v1/silos/${id}` -> show silo details  
- `www.sourcegraph.wtf/v1/silos/${id}/scopes` -> show silo scopes  
- `www.sourcegraph.wtf/v1/forums/${id}` -> show forum details + scopes 

## Views:
Frontend "Frame": Silos (left sidebar), Forums (right sidebar), and Faceted Identities (FIs) (topbar)

1. Document view: Has ability to have n panels, initially just a few.   
    - Read Panel  
    - Write Panel
2. Silo View: 
    - Show constituent scopes/sources

## Directory Structure 
- `containers`: in react containers are components that deal with [how the application works](https://medium.com/@yassimortensen/container-vs-presentational-components-in-react-8eea956e1cea) rather than how state is displayed (component). This container should include the state manager for the documents route. 
- `components`: 
    - `interfaces`: _interprets_ container state, displays it, allows ways of interacting with it.   
    - `reuseable`: individual components that can be reused in different interfaces (e.g. timeline)

## Chunking  
Paging is impossible. Applied or unapplied sources will make the current "focus" position completely irrelevant. Therefore it's out. 

How to request the document at a particular version? 
Op 1: Request each document version - works for text, not for other types. E.g. getting the full version of video for every version would be impossible. Need to apply a patch to existing state. 
Op 2: Request the document version within view

How to move through document versions? 
Op 1: Patching + Anchoring: the frontend needs some general notion for how to patch state and it needs some notion of anchoring to the correct position (some combination of index value and source)

For the v0 implementation get the full documentation each verion request. This should be feasible for text. 

## Separation of Concerns 
Server-Side: Apply transformations, calculate patches between documents 
Client-Side: Apply patches to current state 

Patches are a v1 concern. 

## Routing 
In the documents view I want to manage multiple RW (read/write) panels simultaneously. 

This router will need to be applied only for the documents view (but should save state even when the documents view is left - i.e. the docs and state should be cached/saved).

When a panel is selected a `focused` boolean will be set to `true`. While focused all requests will apply for that scope, as identified by its uuid. This means the router will only be applied to one scope at any given time. 

## Faceted Identities 
All API requests should include an authorization context, which include cookies belonging to each FI.  

Give a UI options for all FIs, explicitly marking them includes the corresponding FI with the request. The cookie is then translated into the corresponding identity in some authn/authz step in the API. That way only the selected FI will be included in the responses.
 
## Timeline 
One frontend feature I would like is he ability for a document to be viewed over its full history. The UI component would be a timeline on the botton of the UI panel, which allows sources to be applied or unapplied.

### Modes 
Read mode should have two modes when moving through versions: "fixed" (based on source or character index) or "following" (show applied sources). This would require a patch-based model, which is probably a v1 concern. 

## Interface for Interfaces 
v1 concern.  

## Login
v1 concern. 

Each FI should require a different secret. Therefore, each FI should initially take an email and password and optionally a cellphone number (for MFA). All information should be encrypted in our database. Over time as people get used to the idea of multiple FIs there must be a more convenient way of handling FIs. I would advise thinking about a lightweight device client. The device client would be like a password manager. Ideally looking into the future I would like to allow yubikeys and public key cryptography, so that should be in the back of our mind. 