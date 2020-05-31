#V0 TODO


## V0 Requirements
1. Text: Only supports text file types.  
2. Local: Only constructs sourcegraph on a single computer. 
3. Bidirectional: Links can go both ways. 
4. Persistent: Links cannot disappear. 
5. Pure Transformations: Transformations should be pure functions to ensure documents are deterministic. 
6. Invertible Transformations: Only if the transformation is inverted can it be unapplied efficiently. 
6. Version-controlled: By default the history of changes is saved.
7. Scoped: Sources can be organized into collections. A document is one type of scope, it is a collection of connected sources. 
8. Set Logic: Set logic can be applied to collections, such that new scopes can be derived from existing scopes.  

## V0 TODO
1. Text Editor React Front End
	- Save button converts text box to document sends to backend
2. Encoding link into file (some kind of markup?)
3. Link Database UUID -> (Source_Head, Source_Tail), Links are functional
4. Git-like diff management for versions, Edits are functional
5. UI for visualizing sourcegraph (to motivate further development on scope UI/logic)


