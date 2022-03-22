# Sourcegraph

## Table of Contents
- [Problem](#problem)
- [Objective](#objective)
- [Philosophy](#philosophy)
- [Description](#description)
- [Transformation Examples](#transformation-examples)
- [Scope Example](#scope-example)
- [Interface](#interface)
- [V0 Requirements](#v0-requirements)
- [V0 TODO](#v0-todo)
- [Data Structure](#data-structure)
- [Social Abstractions](#social-abstractions-siloes-and-forums)
    - [Siloes](#silo-vertical-trust-network)
    - [Forums](#forum-horizontal-collider)
- [Faceted Identity](#faceted-identity)
- [The Place of the Creator](#the-creator-as-the-middleman)
- [The Economic Layer](#the-economic-layer)
    - [Siloes](#siloes)
    - [Forums as Niche Markets](#forum-markets)
- [Layering](#layering)
- [Notes](#scratch-notes)

## Problem

Our mind aggregates and connects sensory and symbolic information in a single embedded entity. There is no software that allows a similar process of accumulation, referencing, and interpretation.

## Objective

Build a Mind Palace: software which allows sources to be accumulated and referenced (and through referencing, reinterpreted infinitely).

## Philosophy

Data comes from “(a thing) given” in Latin. I believe this is subtly incorrect. “Given” implies that reality has faithfully offered the data, and the viewer receives it. It over-emphasizes the place of reality and under-emphasizes the role of the observer. It puts reality in the active role and the observer in the passive role. “(a thing) taken” would be far more appropriate: septa. Observation takes the active role, and subtly distorts the object viewed (as empirically proven by the Heisenburg Principle). The inevitability of distortion should be baked into the language we use, and "data" connotes an arrogant objectivity, where the information collected is reality itself. This is false and misleading. Septa is the imprint of reality that an observer has actively collected and subtly distorted. In Latin septum means "something that encloses". This also seems fitting, the imprint always implies a boundary, it cannot look at everything, it focuses, and focus implies clarity of some things at the expense of loss of others. Focus in bounded. A septum is a bounded imprint of reality.

## Description

A Source is first-degree septum: it is the first observation.

In the future a Source will be an image, video, audio, "sensory" data that is signed upon creation with a hardware provider's private key (or, more likely, a secure hardware enclave on the device, which in turn can prove its hardware provider identity) to prove it has not been distorted or adapted. In an era of deep-fakes hardware providers will be entrusted with the integrity of "sensory" data. We must anticipate and build for that future.

Text is always a nth-degree septum: it is always referring to a first-degree septum (like an image) or to other nth-degree septa (other interpretations). For every "degree" there is another level of re-observation and interpretation, and every re-observation comes both clarifying and distorting simplicity. It is critical to tell which septum is "closer" to reality, the septum that has been observed least. When I reference an image and write about it the edge and the text represent second-degree septa, it is commenting on the image. When I reference that same text to related ideas that is a third-degree septum, etc. For every added degree, the risk of drift and distortion increases, but the upside is clarifying simplicity.

Is a document a single source? A document was created from a constant process of being re-referenced and re-interpreted. A document is a collection of sources with a specific relationship between each source, but that doesn't give it a sacred or special status compared to any other set of references. Therefore, a document is one application of a scope. A scope is a collection of sources. Versions only make sense in the context of scopes. Sources are immutable. Scopes can be taken as an argument but their history is immutable.

What is that relationship in a text document? The process of editing a document takes two forms: writing new content (appended to the end) or editing existing content. In both cases the new content is referencing the current version of the document, and the current version of the document is a chain of sources.

A text document can be imagined as:
```
v5 = f(s4, f(s3, f(s2, f(s1, s0))))
```
The fifth version of the doc is the product of its constituent sources. `f()` expresses how the sources relate to each other, which can be expressed through the `replace` method. While a scope can be a simple collection of sources, in the case of a document it also expresses some relationship between the sources. I suspect (emphasis on suspect) the power of scopes will only be fully realized if we can generalize and express relationships between sources such that different scopes could produce wholly different end-documents in a meaningful way. To see an example of this, jump to the [Scope Example](#scope-example).

There is nothing that forces a document to be a linear document, other than the physical mediums it's based on. Books as a physical medium imposed a linear structure because a graph-document is inconceivable to publish. Linearizing simplified publishing. Speaking (i.e. the source of language) has a vital timing component: two words cannot be spoken at once (and to a lesser extent two thoughts cannot be fully conceived at once), consequently the structure of language is linear (vis-a-vis sentences). This is not true with writing. Reading is time-agnostic, in the sense that _two things can be spoken at once_ in a document. This software will allow the emergence of graph-based documents where narratives can branch and re-emerge seemlessly: braid docs. This is already what fanfiction, comics, research, and the Bible are. They are interweaving stories written by many authors often following (and bounded by) some larger overarching narrative.

Let's imagine a scenario where we have an old-school text document like Lord of the Rings. Single author, no references to external scopes. Now let's imagine a fan-fiction community growing around that document, branching off (and branching back in!) at every conceivable point of the document. For scopes to be powerful they will need to be able to express the base document, the entire braided fan-fiction universe document, and all the sub-story documents in a simple way.

In my current thinking, an end-document is imagined as a chain of transformations of its constituent sources. Philosophically, the "edge" is a functional transformation, applied to the current version of the end-document and a newly created source. This feels philosophically compelling, the process of observation _always_ distorts, and the edge denotes that transformation. This example assumes that a scope is append-only and "moving forward in time". Ideally we will generalize the functional transformations to allow for documents that do not share this same assumption.

If you jump to the [Scope Example](#scope-example) you can see this in action.

A scope must be able to express relationships between sources. It must also be able to express relationships between scopes. Wikipedia is a top-down encyclopedia: it starts with a topic and then fills in the topic with text, images, and external references. Scope-on-scopes could allow bottom-up encyclopedic knowledge simply by taking smaller, more specific scopes and aggregating them into a larger scope. Initially one way of thinking about this is set logic. Can scopes that use other scopes produce end-documents? Should it be allowed to make an end-document from its chapter scopes? This is one example of an end-documents from constituent scopes. I can more easily picture it as an aggregation and categorization tool. This is a known unknown that I note in the [Known Unknown section](#known-unknowns).

Can sources be deleted? Instinctively I am highly highly against source deletion. History feels like something sacred, and it feels highly destructive. When `doc2` references `doc1` it feels like the `doc2` at the very least the source should be completely persistent. I could be convinced that the version of `doc1` used at the time of the reference could be "hidden" but I'm also against that. The idea of the latest version of a document being converted to a private scope seems very reasonable to me, provided all earlier version retain the exact same privacy setting.

##Interface

The interface should deal with two things:
1. Visualizing the accepted format(s)
2. Mapping UI `actions` to functional transformation(s).

An Ipad highlight using an Apple Pencil and a cursor highlight are two distinct UI `actions` but each should make to the same functional transformation for creating a source. The highlight visuals are part of the `action`. Behind the scenes, in both cases a new source is being created from the current scope version.

How mutable is an interface? Should the interface in any way impact the sourcegraph version? What limits what interfaces can view a specific scope end-document? One option is that for a scope to be viewable in an interface, the interface must support _all_ its constituent transformation types. A second option is that the type of the end-document must be supported. A third option is both: an end-document can be supported to the extent defined by its scope.

Naturally the interface types that dominate will accommodate the most transformation types in a cogent way. On the other hand, specialty interfaces would be created, similar to software like photoshop today (where, again, each button represents a transformation on some enclosed source).

What dictates whether two scopes can interact? An overlap in transformations? This is a known unknown that I note in the [Known Unknown section](#known-unknowns).

## Transformation Examples
Transformation | UI | Explanation
1. New source | Highlight | a new source is created from a version of the end-document of a specific scope. Why is a new source necessary? Why not reference the versioned end-document directly? Going back to the philosophy section, the enclosure itself is a transformation that must be made explicit in a source.
2. Reference | Embedded Source | The embedded source would relate to a source that enclosed a specific scope version.
3. Replace | Edit | Document appendments and additions can both be described with the `replace` method
4. Branch | The document _branches_ | In the graph-doc the document splits, which will likely be represented as a bidirectional link  [WIP: Not sure how to think of this, but it seems like an important transformation. Currently seems like it would require the concept of separate scopes to work - this is still assuming the end-document is a text-blob, rather than an actual text graph... food for thought...]
5. Hook | A 

For this idea to work, all file-types must be able to fit into _our_ system, not the other way around.

## Scope Example

Let's imagine a scientific paper. The paper is broken into different chapters, but the edits for each chapters were not made in order; the author went back to edit previous chapters.

Let's imagine there have been 10 edits for 3 chapters:

Chapter 1: `s1`, `s2`, `s7`
Chapter 2: `s3`, `s4`, `s5`, `s9`
Chapter 3: `s6`, `s8`, `s10`

There are two edit types:
Edit (`f()`): `s1`, `s2`, `s3`, `s5`, `s6`, `s8`, `s9`
Reference (`g()`): `s4`, `s7`, `s10`

The document scope can be imagined as a function: `doc_10` = g(`s10`, f(`s9`, f(`s8`, g(`s7`, f(`s6`, f(`s5`, g(`s4`, f(`s3`, f(`s2`, f(`s1`, `0`))))))))))
`Chapter1_3` = g(`s7`, f(`s2`, f(`s1`, `0`)))
`Chapter2_4` = f(`s9`, f(`s5`, g(`s4`, f(`s3`, `0`))))
`Chapter3_3` = g(`s10`, f(`s8`, f(`s6`, `0`)))

The underscored value specifies the version of the scope.

These four scopes should be consistent.

If Chapter3 is edited, how can that transformation be applied to `doc_10` and `Chapter3_3`? For it to be applied, all produced scopes would need to be stored in-memory and applied to each scope. This would likely need to be expressed in a UI option, e.g. a checkbox for current scopes.

If a sentence in Chapter2 is highlighted, creating a new source, only one new source would be created, but it would be applied to all scopes that apply to it. Thus, the transformation is a scope-level concept.

Let's imagine author2 overlaps `s1` to `s7` with their own scope. If author adds `s11` it _could_ be difficult to infer whether the edit is applied to the new scope. Therefore it should be made explicit.

Let's say that the author starts Chapter4. How would they specify the new chapter? There would need to be an accessible UI to create the new scope at the point that `Chapter3_3` is complete.

While implicitly `Chapter3_3` _is_ enclosed by `doc_10` should there be a way to specify a relationship between the scopes? This is a [known unknown](#known-unknowns).

What limits the transformations to `f()` and `g()`? I believe this should be determined at the UI level. It's limited to the number of effective ways the browser can interact with the end-document.

How should signature types (arguments, outputs) be considered? This is a [known unknown](#known-unknowns).

Should an argument only be able to take two arguments, the original document and the new source? Will there ever be a situation where multiple values are required at once? For v0 it may make sense to use two arguments, and see if we run into any snags.

## Use Case: Synthetic Biology [WIP]

How to show a sequence of plasmid transformation as a simulation?

1. There should be a scope specifically for the transformations, and there should be a time dimension in the source-material
2. The history of edits is still saved, but it's separate from the end document.

## How to Make a Graph Doc? [WIP]

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

## Data Structure [WIP]

Each source is a row.

Each scope is a chain of function calls with its constituent sources as arguments.

Should each version of a scope end-document be re-saved? That's what github does. While it works for text and code, it would be prohibitively storage-heavy for images and video. If, however, this service exists in the cloud, perhaps it could be storage-light (i.e. only stored at the source-level) and processing-heavy (i.e. end-documents are produced on the fly). A middle-option is having some notion of check-points.

For transformation-heavy types, maximize the number of checkpoints. For storage-heavy types, minimize the number of checkpoints. This is a fun optimization problem that should absolutely be abstracted away from the user.

I suspect that the optimization would be inclined towards more checkpoints rather than less, from a UX and cost perspective storage is "cheaper" than processing.

```
CREATE TABLE source {
    id uuid NOT NULL, -- immutable 
    degree uint64 NOT NULL, -- immutable
    creator uuid NOT NULL, -- immutable  
    pubkey jsonb NOT NULL, -- 
    signature json NOT NULL, -- non-repudiation, ownership can be tied cryptographically to creator 
    created_at timestamp NOT NULL, -- immutable
    content jsonb NOT NULL, -- immutable 
    parent uuid, -- can be null, null if first-degree septa 
    transformation uuid NOT NULL, -- 
    PRIMARY KEY (id),
}
// Presumably the transformations will be stored in code. We probably will want a transformation-layer where db results are returned and immediately applied in transformations, and then the transformed result is returned to the user 
CREATE TABLE scope_doc {
    id uuid NOT NULL,
    version uint64 NOT NULL,
    doc_parent uuid NOT NULL, -- Can it only reference a single parent scope_doc? Or can multiple be involved? 
    source uuid NOT NULL, -- Ever a situation where multiple sources are applied at once?  
    transformation uuid NOT NULL, -- id for transformation method
}
``` 

### Social Abstractions: Siloes and Forums

It is vitally important we have a clear model for how sourcegraph privacy and collaboration works, specifically how sources and scopes are shared. I think two, and only two, main abstractions are necessary to do this: siloes and forums.

#### Silo: Vertical Trust Network

Siloes own scopes. A silo is a vertical. A vertical is determined by a shared sense of membership and trust. A nuclear family and an extended family are two distinct siloes. A lab is a silo. A college is a silo. A team on a company or the company itself are siloes. Siloes are trust networks. All scopes and sources created for a silo are fully and unconditionally accessible to all members of the silo. From the outside, a silo in today's terms is a _private_ entity, access to its scopes can be sold to outsiders. From the inside, in today's terms is a _public_ entity, all members have complete and unconditional access to those same scopes and those scopes are _public commons_ that must be cultivated. Members will receive shares in the silo that gaurantee them a cut of the total silo income. If a member _corrupts_ the commons, they'll eventually be _exiled_ from the silo. The conditions are implicit: the memberships of the silo itself. If new conditions arise, make a new silo with the appropriate membership and either get permission from the previous silo to fork their scopes or bring the original sources belonging to the members and create a new scope from them. By default there is a global scope where scope and sources are accessible to all user accounts (faceted identities). Faceted identities start with their own personal silo and can create as many silos as they want, which can be used to group scopes. Currently I'm inclined to think that one scope may only belong to one silo, but I could be convinced otherwise.

#### Forum: Horizontal Collider

Forums can share scopes. A forum is a horizontal _between_ verticals, specifically verticals that _depend_ on each other but don't _trust_ each other. When a forum is created, a set of silos is added to the forum, and the silo members have access to shared scopes, but in a conditional way (i.e. access may be contingent on money transaction... more on that later). Forums exist all around us. Markets in both the abstract and real sense are forums, they are places where _verticals_ (producers, buyers, sellers, companies, etc) meet for specific reasons relating to cross dependencies. A career fair is a forum between companies and the college. Calapalooza is a forum between student groups and the college. Enterprise tech conferences are forums between the provider company (Amazon/AWS) and its client companies. Forums can be long-term (like a markerplace) or they can an event (a career fair). So why can't this occur within a silo? Siloes imply trust and therefore access. Forums are places for different siloes to meet and satisfy their mutual needs, not for them to trust each other with full information.

##### Silo Philosophy: Alignment Through Surplus and Pay-It Forward

The intention of silos is to _align_ the incentives of all its members around collective ownership of scopes. Silos are necessary for creating collective cash flow, organizing scopes around relevancy, and enabling privacy. Whatever surpluses that are produced are distributed to the members of the silo. One other _hope_ for silos is to enable "pay-it forward"/karmic economics. This has been a pet idea of mine for a while. A dominant way we currently engage with needs is through transactionality. I pay you X for the Y I need. A necessary middleman for this to work is the state, both for enforcement of the transaction but also for the currency used in the transaction. An _alternative_ model for some needs, in particular for care and education (ie community and creativity), is a "pay-it forward" model. Picture this: I give a friend something - care, money, opportunity, knowledge, a helping hand - but on a loose condition: when given the chance, pay something similar forward to a future person in need. If my friend faithfully does this, and the recipient of their care does the same thing, and all the future people in the chain of their care do the same, my one action has turned into an _infinite_ chain of care. And some of that care will come back to me, indirectly. This is the logic behind karma. _The problem_ with this logic is that the original "igniter" of the chain of care will _only_ do so if they have the security to do so. The best way of gauranteeing that security is if the "chain of care" comes back to them eventually, justifying their original "gift" and giving them the security (and faith in the logic) _to give more_. The smaller the silo, the more likely that chain of care will come back to its original gifter, enabling the chain in the first place. Therefore silos could be an ideal environment for "pay-it forward" logic. The upside of the logic is drastically reduced transaction costs around a variety of needs. Today, we need to pay some for everything we need. This seems there are transaction costs around _everything_ we need. In a world of pay-it forward logic the transaction costs _within_ a silo are zero.

#### Forum Philosophy: Cosmic Serendipity

Cosmic serendipity is an idea I've been thinking about for years. Cosmic serendipity is the logic behind fortuitous collisions occurring. When I run into my neighbor in an international airport, this is cosmic serendipity at work. You would think that in a world of 9 billion people the odds of running into a neighbor in an international airport around the world is unlikely. But this is assuming that there is no underlying structure behind _who_ will be at that airport. In point of fact, there _is_ structure and there's a simple probabilistic argument that expresses its default behavior: the Birthday Paradox. If you have 23 people in a room what are the odds of two people sharing the same birthday? Given there are 365 days in year, and 365 appears to dwarf 23, the initial intuition are low odds. What does your intuition tell you?

In fact there is a 50% chance of a birthday “collision”, of two people sharing a birthday. How can that be explained? There are 365 days but the number of people is not the important value to consider, it’s the number of relationships between those people. In the group, the first person has 22 other relationships. The second has 21 unique relationships. The third 20. Etc. This summation equals 253 unique pair-wise relationships. There is a 1/365 chance that two birthdays “collide”, which means there is a 1 - (1/365) chance they do not collide. The trick of the birthday paradox is looking for the odds that zero relationships collide, which are the odds that no collisions occur over all unique relationships: (1 - (1/365))^253 = 0.499488. Now, it’s easy to identify the odds of one or more collisions, because it’s everything else, it’s 1 - 0.499… = 0.50. By this same process, 50 people: 97% chance of collision. 75 people: 99.95%.

Life is an unconstrained Birthday Paradox. You aren’t just talking about birthdays colliding, you’re talking about everything colliding: shared experiences, values and ideas; chance encounters with friends in cities, international airports, transportation hubs; strange number collisions; collisions upon collisions upon collisions. And if you begin to see the world in these terms, you can begin playing the odds.

Models are mathematical stories, and what this story demonstrates is that there are two important factors to consider where coincidence is concerned: the odds of collision and the number of unique relationships for a particular kind of collision. Even if the odds of collision are seemingly very small (1/365 for a birthday), if they’re applied to many relationships the odds of occurrence become extremely high. Put another way, connectedness (number of unique relationships) exponentially varies with coincidence (occurrence of a collision). _But_ if the odds of collision _are also_ high, then the odds of cosmic serendipity become increasingly certain.

Note that in your own life, the loci of cosmic serendipity are "horizontals": transportation (buses, subways, ubers, planes), conferences, bars, restaurants, theaters, parks. All these venues cut across many separate/non-interacting verticals, and that's what makes them so vitally important. This is why forums are horizontals between silos. The difference, however, is that they are _explicitly_ intended for cosmic serendipity and by _using_ silos will be highly "filtered", _ideal_ environments for fortuitous collisions.

### Faceted Identity

Identity is a "public" (public _within_ a specific community) history of one's actions. My identity in one community (my family) is completely different from other (my company). This is because the public record of my actions differ between these communities. In other words, the only person with a "complete" idea of my identity, my history of actions, is me, and different "facets" of that identity are presented based on selective knowledge of that full history of actions, often at the level of the community.

One of the reasons online servives are so annoying to use is that there's no way to control these different personas, these faceted identities, between communities, leading to awkward contradictions. Instead this reality should be made explicit through the abstractions around identity we use, we should _allow_ people to be multiple separate identities that they can closely control. In this world, _information absence_ is the default, rather than presence. By default I will _only_ know information about faceted identities from the silos (and therefore forums) that we share. This will prevent awkward "bleeding" of information between different communities.

The user's client will be responsible for managing those identities (initially through cookies, later through an actual client which manages public keys).

A user should also be allowed to _present_ these faceted identities and their memberships to outsiders. Resumes today are essentially a list of our memberships along with some metrics (income, awards, grades). Presentation of selective identities and/or their silos could be used for the same purpose.

### The Creator as the Middleman

The creator owns their sources and can prove their ownership. Creators must be given rights to their sources. This ultimately is what gives bite to individuals within a silo: the work they have contributed. This means that individuals can always control content they create, and can extract rent/leverage from that content.

That being said, if a second user made a source from a now-closed source, these sources will be respected. An owner can choose the current state of a source/scope, but they cannot discount the right others have to keep their sources.

This approach seems like the best way of balancing the rights of creators and users. It gives creator's leverage while also respecting the rights of others.

### The Economic Layer

#### Siloes

Silos can be imagined as both a private _or_ public entity in today's world. While they're a commons of information, every member of a silo can receive a share of its total income. The members of the silo itself should be able to determine how their shares are controlled.

#### Forum Markets

Forums should also be imagined as niche marketplaces where consumers pay creators to be able to (a) view scope end-documents and (b) use sources/scopes for their own content creation. If there are cross-dependencies between a user and a silo it creates the incentive for the user to join the silo annulling costs. In other words, forums are a necessary component for the emergence of new silos - cross-dependencies beget new silos.

### Layering
What is the relationship between sourcegraph, social abstractions, and the economic layer? How can they be cleanly organized?

## Scratch Notes:
Imagine a world where new end-documents could be created between an infinite number of file types, creating currently inconceivable document formats and creative production. I want a Renaissance in the framing of thought itself, Sourcegraph is a meta-framing, and its strength will come from how expressively it describe new framings.   \
-- privacy information should be immutable, set at the time of creation and respected forever afterwards... but that also means it can't be publicized... Seems like a one-direction issue, i.e. private -> public is possible but the reverse is not.
- reference, makes new source, most basic transformation is enclosure
  Does it require a defined transformation? Is there a case where a source should be created but the transformation not recorded? Maybe at the base-layer all transformations to produce new sources are created. Let's go
  Where do edges fit into this model? If a septum encloses, it must enclose something. The edge points to what is being enclosed and distilled. Therefore, the edge is fundamental to what the septum is describing. A first-degree septum has no edge: it connects to Reality.
  I'm inclined to think of an edge as fundamental to a source. What does it even mean to have an edge between two undefined documents? An edge only makes sense within its necessary context.
  The downside of this approach is that an individual can take a silo hostage, if they've contributed vital sources. While this is true, an individual would be sacrificing their future credibility as well, indicating either extreme integrity or desperation, either way such extremes indicate a real underlying problem, and giving recourse to individuals is essential for solving problems (or, even better, preventing them in the first place).
- Increasingly I think a source and edge are fundamentally related. Is there any velocity-position metaphor that can be drawn? Intuitively they seem related (and I would love to rope the Heisenberg Principle into this idea)