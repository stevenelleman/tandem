__Sourcegraph__

# Table of Contents
* [Problem](#problem)
* [Objective](#objective)
* [Note](#note)
* [Feature List](#feature-list)
   * [Version 0.1](#version-0.1)
     * [Later Versions](#later-versions)
   * [Aspirational](#aspirational)
* [Philosophy](#philosophy)
* [Feature Description](#feature-description)
   * [Interface](#interface)
   * [Client Application](#client-application)
   * [Transformation](#transformation)
   * [Plaintext](#plaintext)
   * [Scopes](#scopes)
   * [Arrows](#arrows)
   * [References](#references)
   * [Time Dimensions](#time-dimensions)
   * [Types](#types)
   * [Views](#views)
   * [Type Attributes](#type-attributes)
   * [Custom Transformations](#custom-transformations)
   * [Tiles](#tiles)
   * [Sensible Search](#sensible-search)
   * [Creation Suites](#creation-suites)
   * [Attribute Axis](#attribute-axis)
   * [Universal Axis](#universal-axis)
   * [Universal Library](#universal-library)
   * [Monetization Scheme](#monetization-scheme)
   * [Privacy Features](#privacy-features)
   * [Multi-Device Support](#multi-device-support)
   * [Type Execution](#type-execution)
   * [Anchor Time Travel](#anchor-time-travel)
* [Move to a Different File](#the-following-should-probably-be-moved-to-a-different-file)


# Problem

Our mind aggregates and connects sensory and symbolic information in a single embedded entity. There is no software that allows a similar process of accumulation, referencing, and interpretation.

# Objective

Software that allows users to collect, reference, interpret, and share sources.

## Note

A **sourcegraph** (lowercase s) is a group of documents organized with connections and scopes. It's defined using a series of transformations (deltas) on another sourcegraph (originally, an empty one).
**Sourcegraph** (capital S) is the software that allows users to make sourcegraphs.

## Feature List



### Version 0.1 

* **Interface:** Make one that is good. TODO: MAKE THIS SUMMARY BETTER.
* **Client Application:** Sourcegraph runs and stores all data on the user's local machine using cloud technologies.
* **Transformation:** A sourcegraph is defined as a series of git-like functional transformations applied on a blank sourcegraph. A transformation is a pure function to ensure documents are deterministic.
* **Plaintext:** Users can create and edit plaintext documents
* **Scopes:** Users can draw scopes around documents to represent categories they belong to. Scope s are infinitely recursive (scopes within scopes within scopes...) and their contents disappear when the user "zooms out" enough.
   *  **Set Logic:** Set logic can be applied to scopes to support **sensible search** and so that new scopes can be derived from existing scopes..
   * **Arrows:**  Users can draw arrows between documents and label the arrows to show what relationships they represent.
   * **References:** Sourcegraph keeps track of where the user gets their ideas from by displaying relationships between referencing and referenced document. These relationships persist even if the referencing snippet is deleted.


### Later Versions

* **Time Dimension:** Each sourcegraph is a series of transformation layered like sediment on an original sourcegraph. Users can see these transformations as a **chain** of *deltas* "changes" and "look back" to what the sourcegraph looked like on an earlier delta. They can also **branch** away from the chain at a point to make an "alternate version" of their sourcegraph. They may also be able to **merge** two sourcegraph branches together if desired.
* **More and Custom Types:** Sourcegraph supports different *types* of documents ranging from musical scores to video. Users can create their own types and define transformations on them.
* **Views:** Types have their own UI or *view*, which can feature transformations that the view creator deemed important for that type.  A view is like a workspace for a type.
* **Type Attributes:** Different types of documents will have different attributes "author, date created, tempo etc." that the user can fill in. Users can define their own custom attributes for custom types or modify already-existing types with new attributes. 
* **Custom Transformations:** Users can make their own transformation types that take an arbitrary number of parameters can apply to individual documents or the sourcegraph itself. Users can customize their UI to feature their favorite transformations.
* **Tiles:** Users can see multiple views at once with a tiling window manager. Tiles "look at" the sourcegraph in ways the user can define and arrange themselves in sensible ways. The user can cluster and uncluster tiles into tabs, and they can save their favorite clusters. (One tile, one view) 
* **Sensible Search:** Combines regular expressions and scope set logic for a powerful, elegant search tool.
   *   A **collapsing reference** happens when you *collapse* part of the major document into a minor document in a way that still allows you to see the contents of the minor document in their proper place in the major document with a  pres of a button.


### Aspirational

* **Creation Suites:** User-made creation tools built on Sourcegraph's UI logic. Users can make music, art, and poetry from within Sourcegraph with high-quality interfaces that rival industry standards like the Adobe Creative Suite.
* **Attribute Axis:** For continuous numeric attributes, users can sort documents by attribute to create a chart. (For example, the user could sort documents with the 'city" attribute by Founding Date and Average Temperature if city documents had both these attributes. ~~Note, we need to consider what happens when a document has a blank attribute. If the user forgot to fill in a founding date for Amsterdam, what happens when they sort cities by founding date?~~)
* **Universal Library**: Users can share documents, types, and transformations they've made on a paid or unpaid basis with other users.
* **Monetization Scheme:** Users can make money with their creations on sourcegraph by charging other users a small fee to use their content, types, or transformations. This wouldn't be a one-time fee but rather a continuous cash flow that increased in proportion to how much value a downstream user got from the upstream user's content. Creators would be incentivized to make content that is high-quality and easy to remix, because the more users "branch off" from their content and become popular, the more money flows to the original creator. At the same time, there would also be a "distance discounting" feature so that the majority of the cash wouldn't flow to creators in the "far past." This will incentivize even popular creators to keep making new content.
* **Privacy Features:** Users can define who can view, edit, and administrate scopes through the use of bounded communities called **silos.** Silos can pool creative resources and distribute income to members. In **forums**, silos come together to trade creative knowledge for money.
* **Multi-Device Support**: Sourcegraph should support mobile and VR devices as well as Windows, Linux, and Mac. 
* **Type Execution:** Sourcegraph should support creators *executing* documents of varying types from within Sourcegraph. For musicians, this means being able to play the songs in Sourcegraph. For programmers, it means compiler support. For animators, it means the ability to pause and play animations. 
* **Anchor Time Travel**: Users can highlight, or "anchor" to a certain piece of text or document and move back in time specifically for that text to see how it has changed.


## Philosophy
                
*Data* is Latin for “a thing given”. We believe this is subtly incorrect. “Given” implies that reality has faithfully offered the data, and the viewer receives it. It over-emphasizes the place of reality and underemphasizes the observer's role. In this document, we will prefer *septa:* “a thing taken.” This emphasizes the role of the active observer, who distorts the viewed object. ~~Which dictionary did you get septa from? I only see it in wiktionary as the plural of the neuter noun *septum*, "enclosure." Datum is a common present participle of *do* "I give." I don't know of or see any Latin word for "to take" from which we may in like manner derive *septa.* I could be messing something.~~

Sourcegraph has three ideals: *persistence*, *user ownership*, and *version control.*

### Persistence

Can sources be deleted? Instinctively I am highly against source deletion. History feels like something sacred, and it feels highly destructive. When `doc2` references `doc1` it feels like the `doc2` at the very least the source should be completely persistent. I could be convinced that the version of `doc1` used at the time of the reference could be "hidden" but I'm also against that. The idea of the latest version of a document being converted to a private scope seems very reasonable to me, provided all earlier version retain the exact same privacy setting.

### User Ownership

The user will be able to own almost every aspect of sourcegraph. Firstly, the sourcegraph code itself will be open-source, so tech-savvy users can make their own splinters of the software. However, most user control will be built into the software itself. Content, transformations, document types, and interfaces can all be generated, shared, and monetized by users.

### Version Control

Much like git, Sourcegraph allows users to view and manage the history of their sourcegraphs. A sourcegraph is defined as a series of transformations, and the user can view the entire chain of transformations back from its origin point. They can also view the state of the sourcegraph at an earlier transformation and, if they desire, fork the sourcegraph from that point in order to make an alternate sourcegraph.

Later, when  users are able to share sourcegraphs with each other, the Sourcegraph ecosystem will be based largely on following and forking sourcegraphs in silos.

Let's imagine a scenario where we have an old-school text document like Lord of the Rings. Single author, no references to external scopes. Now let's imagine a fan-fiction community growing around that document, branching off (and branching back in!) at every conceivable point of the document. For scopes to be powerful they will need to be able to express the base document, the entire braided fan-fiction universe document, and all the sub-story documents in a simple way

## Feature Description


 **Version 0.1**
---


### Interface
                
The interface should deal with two things:
1. Visualizing the accepted format(s)
2. Mapping UI `actions` to functional transformation(s).

On an iPad, you can highlight text with an Apple Pencil. On a laptop, you do that with a cursor. These are two distinct UI `actions`that map to the same functional transformation. The highlight visuals are part of the `action`. Behind the scenes, in both cases a new sourcegraph is created by applying a transformation to the current sourcegraph.

How mutable is an interface? Should the interface in any way impact the sourcegraph version? What limits what interfaces can view a specific scope end-document? One option is that for a scope to be viewable in an interface, the interface must support _all_ its constituent transformation types. A second option is that the type of the end-document must be supported. A third option is both: an end-document can be supported to the extent defined by its scope.

The dominant interfaces will incorporate a wide range of the transformations that the users like the most. On the other hand, users will create specialty interfaces, similar to software like Photoshop today (see [Creation Suites](#creation-suites)).

### Client Application

Version 0.1 will edit and save sourcegraphs locally. However, it will leverage React cloud technologies to easily integrate it in the cloud in the future.

TODO: DISCUSS AND WRITE DOWN IMPLEMENTATION DETAILS, MENTION TANDEM.

### Transformations

In the purest terms, a transformation is a **function** that takes **arbitrary parameters** and changes a sourcegraph.

__Examples of Transformations:__

* Drawing an arrow from one document to another.
* Changing text in a document
* Drawing a scope around three specific documents (in other words, making a new scope and putting three documents in it)

__Philosophy of Transformations:__

A Source is a piece of sensory data (image, video, audio) that is signed upon creation with a hardware provider's private key (or, more likely, a secure hardware enclave on the device, which in turn can prove its hardware provider identity) to prove it has not been distorted or adapted. In an era of deep-fakes, hardware providers will be entrusted with the integrity of "sensory" data. We must anticipate and build for that future.

Text is always a nth-degree septum: it refers to a first-degree septum (like an image) or to other nth-degree septa (other interpretations). For every "degree" there is another level of re-observation and interpretation, and every re-observation comes both clarifying and distorting simplicity. It is critical to tell which septum is "closer" to reality, the septum that has been observed least. When I reference an image and write about it the edge and the text represent second-degree septa, it is commenting on the image. When I reference that same text to related ideas that is a third-degree septum, etc. For every added degree, the risk of drift and distortion increases, but the upside is clarifying simplicity.

Is a document a single source? A document was created from a constant process of being re-referenced and re-interpreted. A document is a collection of sources with a specific relationship between each source, but that doesn't give it a sacred or special status compared to any other set of references. Therefore, a document is one application of a scope. A scope is a collection of documents. Versions only make sense in the context of scopes. Sources are immutable. Scopes can be taken as an argument but their history is immutable.

What is that relationship in a text document? The process of editing a document takes two forms: writing new content (appended to the end) or editing existing content. In both cases the new content is referencing the current version of the document, and the current version of the document is a chain of sources.
A text document can be imagined as:

```
v5 = f(s4, f(s3, f(s2, f(s1, s0))))
```
    
The fifth version of the doc is the product of its constituent sources. `f()` expresses how the sources relate to each other, which can be expressed through the `replace` method. While a scope can be a simple collection of sources, in the case of a document it also expresses some relationship between the sources. I suspect (emphasis on suspect) the power of scopes will only be fully realized if we can generalize and express relationships between sources such that different scopes could produce wholly different end-documents in a meaningful way.

Books are linear because a graph-document is inconceivable to publish. Speaking (i.e. the source of language) has a vital timing component: two words cannot be spoken at once (and to a lesser extent two thoughts cannot be fully conceived at once), consequently the structure of language is linear (vis-a-vis sentences). This is not true with writing. Reading is time-agnostic, in the sense that _two things can be spoken at once_ in a document. This software will allow the emergence of graph-based documents where narratives can branch and re-emerge seemlessly: braid docs. This is already what fanfiction, comics, research, and the Bible are. They are interweaving stories written by many authors often following (and bounded by) some larger overarching narrative.

__Example:__

Let's imagine a scientific paper. The paper is broken into different chapters, but the edits for each chapters were not made in order; the author went back to edit previous chapters.

Let's imagine there have been 10 edits for 3 chapters:

```

Chapter 1: `s1`, `s2`, `s7`
Chapter 2: `s3`, `s4`, `s5`, `s9`
Chapter 3: `s6`, `s8`, `s10`

There are two edit types:
Edit (`f()`): `s1`, `s2`, `s3`, `s5`, `s6`, `s8`, `s9`
Reference (`g()`): `s4`, `s7`, `s10`

The sourcegraph for this paper can be imagined as a function: `doc_10` = g(`s10`, f(`s9`, f(`s8`, g(`s7`, f(`s6`, f(`s5`, g(`s4`, f(`s3`, f(`s2`, f(`s1`, `0`))))))))))
`Chapter1_3` = g(`s7`, f(`s2`, f(`s1`, `0`)))
`Chapter2_4` = f(`s9`, f(`s5`, g(`s4`, f(`s3`, `0`))))
`Chapter3_3` = g(`s10`, f(`s8`, f(`s6`, `0`)))


```

Other transformation ideas: 
- Hook/Flag 
- Todo 

For this idea to work, all file-types must be able to fit into _our_ system, not the other way around.

#### Deletion

**Deletion** is also a transformation, but because you can go back and recover an earlier version of your sourcegraph, deletion is never permanent. It is a transformation more akin to covering a picture with white paint than burning it. Just like you can always scrape off the paint, you can always revert a sourcegraph to an earlier point when the thing you deleted existed.

We imagine that most deletions happen because of mistakes. Because storage is cheap, the user has no need to delete something they can hide instead. Sourcegraph's strong privacy features will ensure that they can.  **In sourcegaph, there is no deletion. There is only privacy.** 

### Plaintext

Sourcegraph will have the ability to edit plaintext. Maybe we can add a markdown type in version 0.2 to support markdown, but for the first prototype, plaintext is enough.

### Scopes

A scope is a collection of documents. A sourcegraph can hold any number of documents and any number of scopes; a single document can be in any number of scopes and a single scope can contain any number of documents. The user can name scopes.

In practice, the use of scopes will vary from user to user. One usecase is to represent *groups* of documents that share a certain attribute that the user thinks is important.

The user can search for documents in any combination of scopes using set logic. For example:

`[SCOPE1]&&[![SCOPE2]]`

Will return all the documents that are in set 1 and that *aren't* in set 2. Syntax details may differ in the final  product.

The underscored value specifies the version of the scope.

These four scopes should be consistent.

### Arrows

Users often want to understand the relationships between documents or objects. To support this, sourcegraph allows them to draw arrows between documents and label the arrows to show what relationships they represent. This will be useful to indicate relationships such as causality, influence, social connection, and chronology. Arrows will be unlabelled by default, the user can click on them to add a label.

Arrows can be unidirectional or bidirectional.

### References

References are a special type of relationship between documents. They are represented by an arrow labelled "references" and have special properties. References are asymmetrical: there is a *referencing* document and a *referenced* document. To make a reference, the user will highlight a snippet of the referenced document and drag it over to a place in the referencing document. The will now be an arrow pointing to the referenced document from the referencing document, and when the user clicks that arrow, sourcegraph will show them the exact snippet the referencing document references.

A **reference** is a superset of the more traditional idea of a link. It's like a link in that if you click on a snippet incorporated in a document, you can be redirected to to the document the snippet came from. It's *unlike* a link in that, even if you delete all the text of the entire snippet, the reference persists. *The text might be gone, but the influence remains.*

### Time Dimension

---
**Later Versions**
---

Later on, the user will be able to view previous transformations of their sourcegraph as a chain of transformations, like commits in GitHub. They can also view earlier transformations and branch a new sourcegraph from one of them, thus creating two alternate versions of the same sourcegraphs. Depending on whether we can easily merge sourcegraphs (we suspect we can), the user will be able to merge these two branches easily.

### More and Custom Types

We'll add support for many more types of documents, ranging from rich text documents to midi files to pngs. We'll give users the opportunity to define their own types and the best transformations on them, so that (ideally) the users will do the work of the developers.

#### Type Attributes

Types will contain many attributes. For example, a user-defined "city" data type might have average climate and altitude as attributes. These attributes can be used as aids for the user and as inputs for  user-defined and already-made functions.

*Question:* *Will the user need to know how to code in order to create custom transformations and types? If so, in what language?*

### Custom Transformations

Just as the user can define custom types, they can define custom transformations on those types. It's important that these two features are released at the same type, because custom types aren't really that useful until the user can transform them easily. These transformations will come with UI elements the user can also define.

*Question:* *Will the user need to know how to code in order to create custom transformations and types? If so, in what language?*

### Tile View

A view is a UI interface with various buttons that activate various transformations. The user can modify views just as they can modify the transformations they contain. If they like, they can see multiple views at once through a type of tiling window manager. This will be supported on laptops, mobile, and VR.

### Creation Suites

Eventually, users will create views that utilize custom types and UIs that are so powerful they rival already-existing creative tools like the Adobe Creative Suite. Though these tools might not have the raw power and capability of already-existing tools, their integration within Sourcegraph, which creators will already be using to organize their sources, might give them an edge.

### Attribute Axis

Users can sort documents in their sourcegraph by a certain numerical attribute. This will make it easier to visualize timelines and chronologies.

### Universal Library

Users can draw from a giant library of documents common to everyone that they would not need to pay for. Users could also add to this repository.

### Monetization Scheme

I'm not touching this too closely for now. Basically, if they want, creators will be able to monetize the stuff that create, whether that be content (video, audio, text, etc.), transformations, types, or views.

#### Silos

Silos can be imagined as both a private _or_ public entity in today's world. While they're a commons of information, every member of a silo can receive a share of its total income. The members of the silo itself should be able to determine how their shares are controlled.

#### Forum Markets

Forums should also be imagined as niche marketplaces where consumers pay creators to be able to (a) view scope end-documents and (b) use sources/scopes for their own content creation. If there are cross-dependencies between a user and a silo it creates the incentive for the user to join the silo annulling costs. In other words, forums are a necessary component for the emergence of new silos - cross-dependencies beget new silos.

### Layering

What is the relationship between sourcegraph, social abstractions, and the economic layer? How can they be cleanly organized?

### Privacy Features

#### Social Abstractions: Siloes and Forums

It is vitally important we have a clear model for how sourcegraph privacy and collaboration works, specifically how sources and scopes are shared. Two abstractions are necessary: silos and forums.

##### Silo: Vertical Trust Network

Silos own scopes. A silo is a vertical. A vertical is determined by a shared sense of membership and trust. A nuclear family and an extended family are two distinct silos. A lab is a silo. A college is a silo. A team on a company or the company itself are silos. Silos are trust networks. All scopes and sources created for a silo are fully and unconditionally accessible to all members of the silo. ~~I don't think scope access should be an all-or-nothing proposition. Read, write, and admin privileges should be separate, like they are in most social media fora today.~~ From the outside, a silo in today's terms is a _private_ entity, access to its scopes can be sold to outsiders. From the inside, in today's terms is a _public_ entity, all members have complete and unconditional access to those same scopes and those scopes are _public commons_ that must be cultivated. Members will receive shares in the silo that guarantee them a cut of the total silo income. If a member _corrupts_ the commons, they'll eventually be _exiled_ from the silo. The conditions are implicit: the memberships of the silo itself. If new conditions arise, make a new silo with the appropriate membership and either get permission from the previous silo to fork their scopes or bring the original sources belonging to the members and create a new scope from them. By default there is a global scope where scope and sources are accessible to all user accounts (faceted identities). Faceted identities start with their own personal silo and can create as many silos as they want, which can be used to group scopes. Currently I'm inclined to think that one scope may only belong to one silo, but I could be convinced otherwise.

##### Forum: Horizontal Collider

Forums can share scopes. A forum is a horizontal _between_ verticals, specifically verticals that _depend_ on each other but don't _trust_ each other. When a forum is created, a set of silos is added to the forum, and the silo members have access to shared scopes, but in a conditional way (i.e. access may be contingent on money transaction... more on that later). Forums exist all around us. Markets in both the abstract and real sense are forums, they are places where _verticals_ (producers, buyers, sellers, companies, etc) meet for specific reasons relating to cross dependencies. A career fair is a forum between companies and the college. Calapalooza is a forum between student groups and the college. Enterprise tech conferences are forums between the provider company (Amazon/AWS) and its client companies. Forums can be long-term (like a markerplace) or they can an event (a career fair). So why can't this occur within a silo? Siloes imply trust and therefore access. Forums are places for different siloes to meet and satisfy their mutual needs, not for them to trust each other with full information.

###### Silo Philosophy: Alignment Through Surplus and Pay-It Forward

The intention of silos is to _align_ the incentives of all its members around collective ownership of scopes. Silos are necessary for creating collective cash flow, organizing scopes around relevancy, and enabling privacy. Whatever surpluses that are produced are distributed to the members of the silo. One other _hope_ for silos is to enable "pay-it forward"/karmic economics. This has been a pet idea of mine for a while. A dominant way we currently engage with needs is through transactionality. I pay you X for the Y I need. A necessary middleman for this to work is the state, both for enforcement of the transaction but also for the currency used in the transaction. An _alternative_ model for some needs, in particular for care and education (ie community and creativity), is a "pay-it forward" model. Picture this: I give a friend something - care, money, opportunity, knowledge, a helping hand - but on a loose condition: when given the chance, pay something similar forward to a future person in need. If my friend faithfully does this, and the recipient of their care does the same thing, and all the future people in the chain of their care do the same, my one action has turned into an _infinite_ chain of care. And some of that care will come back to me, indirectly. This is the logic behind karma. _The problem_ with this logic is that the original "igniter" of the chain of care will _only_ do so if they have the security to do so. The best way of guaranteeing that security is if the "chain of care" comes back to them eventually, justifying their original "gift" and giving them the security (and faith in the logic) _to give more_. The smaller the silo, the more likely that chain of care will come back to its original gifter, enabling the chain in the first place. Therefore silos could be an ideal environment for "pay-it forward" logic. The upside of the logic is drastically reduced transaction costs around a variety of needs. Today, we need to pay some for everything we need. This seems there are transaction costs around _everything_ we need. In a world of pay-it forward logic the transaction costs _within_ a silo are zero.

##### Forum Philosophy: Cosmic Serendipity

Cosmic serendipity is the logic behind fortuitous collisions occurring. When I run into my neighbor in an international airport, this is cosmic serendipity at work. You would think that in a world of 9 billion people the odds of running into a neighbor in an international airport around the world is unlikely. But this is assuming that there is no underlying structure behind _who_ will be at that airport. 

Life is full of such "coincidences:" shared experiences, values and ideas; chance encounters with friends in cities, international airports, transportation hubs.

Models are mathematical stories, and what this story demonstrates is that there are two important factors to consider where coincidence is concerned: the odds of collision and the number of unique relationships for a particular kind of collision. Even if the odds of collision are seemingly very small, if they’re applied to many relationships the odds of occurrence become extremely high. Put another way, connectedness (number of unique relationships) exponentially varies with coincidence (occurrence of a collision). _But_ if the odds of collision _are also_ high, then the odds of cosmic serendipity become increasingly certain.

Note that in your own life, the loci of cosmic serendipity are "horizontals": transportation (buses, subways, ubers, planes), conferences, bars, restaurants, theaters, parks. All these venues cut across many separate/non-interacting verticals, and that's what makes them so vitally important. This is why forums are horizontals between silos. The difference, however, is that they are _explicitly_ intended for cosmic serendipity and by _using_ silos will be highly "filtered", _ideal_ environments for fortuitous collisions.

#### Faceted Identity

Identity is a "public" (public _within_ a specific community) history of one's actions. My identity in one community (my family) is completely different from other (my company). This is because the public record of my actions differ between these communities. In other words, the only person with a "complete" idea of my identity, my history of actions, is me, and different "facets" of that identity are presented based on selective knowledge of that full history of actions, often at the level of the community.

One of the reasons online services are so annoying to use is that there's no way to control these different personas, these faceted identities, between communities, leading to awkward contradictions. Instead this reality should be made explicit through the abstractions around identity we use, we should _allow_ people to be multiple separate identities that they can closely control. In this world, _information absence_ is the default, rather than presence. By default I will _only_ know information about faceted identities from the silos (and therefore forums) that we share. This will prevent awkward "bleeding" of information between different communities.

The user's client will be responsible for managing those identities (initially through cookies, later through an actual client which manages public keys).

A user should also be allowed to _present_ these faceted identities and their memberships to outsiders. Resumes today are essentially a list of our memberships along with some metrics (income, awards, grades). Presentation of selective identities and/or their silos could be used for the same purpose.

#### The Creator as the Middleman

The creator owns their sources and can prove their ownership. Creators must be given rights to their sources. This ultimately is what gives bite to individuals within a silo: the work they have contributed. This means that individuals can always control content they create, and can extract rent/leverage from that content.

That being said, if a second user made a source from a now-closed source, these sources will be respected. An owner can choose the current state of a source/scope, but they cannot discount the right others have to keep their sources.

This approach seems like the best way of balancing the rights of creators and users. It gives creator's leverage while also respecting the rights of others.

### Multi-Device Support

Sourcegraph should support mobile and VR devices as well as Windows, Linux, and Mac. Ideally, the ecosystem for types, transformations, views, and content will be the same for all these platforms

### Anchor Time Travel:
    .
Users can highlight, or "anchor" to a certain piece of text or document and move back in time specifically for that text to see how it has changed.

## Implementation

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




# The Following Should Probably be Moved to a Different File

## V0 TODO

1. Text Editor React Front End
- Save button converts text box to document sends to backend
2. Encoding link into file (some kind of markup?)
3. Link Database UUID -> (Source_Head, Source_Tail), Links are functional
4. Git-like diff management for versions, Edits are functional
5. UI for visualizing sourcegraph (to motivate further development on scope UI/logic)

## Scratch Notes:

Imagine a world where new end-documents could be created between an infinite number of file types, creating currently inconceivable document formats and creative production. I want a Renaissance in the framing of thought itself, Sourcegraph is a meta-framing, and its strength will come from how expressively it describe new framings.   \
-- privacy information should be immutable, set at the time of creation and respected forever afterwards... but that also means it can't be publicized... Seems like a one-direction issue, i.e. private -> public is possible but the reverse is not.
- reference, makes new source, most basic transformation is enclosure
Does it require a defined transformation? Is there a case where a source should be created but the transformation not recorded? Maybe at the base-layer all transformations to produce new sources are created. Let's go
Where do edges fit into this model? If a septum encloses, it must enclose something. The edge points to what is being enclosed and distilled. Therefore, the edge is fundamental to what the septum is describing. A first-degree septum has no edge: it connects to Reality.
I'm inclined to think of an edge as fundamental to a source. What does it even mean to have an edge between two undefined documents? An edge only makes sense within its necessary context.
The downside of this approach is that an individual can take a silo hostage, if they've contributed vital sources. While this is true, an individual would be sacrificing their future credibility as well, indicating either extreme integrity or desperation, either way such extremes indicate a real underlying problem, and giving recourse to individuals is essential for solving problems (or, even better, preventing them in the first place).
- Increasingly I think a source and edge are fundamentally related. Is there any velocity-position metaphor that can be drawn? Intuitively they seem related (and I would love to rope the Heisenberg Principle into this idea)

## General Questions


### Save Automatically?
### Further Discuss Privacy
### Checkpoints
* Why have them if every transofrmation can be reverted back to?
### Functional Invertibility?


### What Did Steven Mean by these?

-   Prescriptive vs Nonsense
-   Character vs Letters
-   Entrenchment vs Security



## Discussions to Have

* Entrenchment vs. security in the context of our monetization system
   * Discounted cash flow could make retiring on sourcegraph income impossible. Is this what we want?


## Sean's Questions

* **Note:** My general philosophy while editing this document was **change, don't question.** If I saw something Steven wrote that I didn't understand, I replaced it with the concept I thought it mapped to in our discussions over the past few days that I did understand. I know that this slash-and-burn style editing could lead to my missing something, but I didn't want to constantly bug Steven with questions while he was working and figured that we could revert any unwise deletions. In the cases where I did choose to question, I put these questions here.

22.  Going meta:
23.  Transformation, Action, Type, and View are also types
24.  Can read/write them with their own T / A / V
25.  Represent default types
26.  For this to work need to define simple contracts between each component
27.  Users may specify the version of the T / A / V as they change (although it feels unlike T and A will change too much)
28.

*I'm skeptical of this.* If the code that drives types is also a Type that the user can modify, then the user can break Types. If they break types, they break everything that uses Types, which is their entire Sourcegraph. If the time dimensions also uses types, they break that too and can't branch off from an earlier version to revert their changes. In summary, this could allow people to irreparably ruin their sourcegraphs. 

* Why is a source signed upon creation with a hardware provider's public key? 


* I can see the rationale behind having transformations be pure functions, but will this become too restrictive?

From **transformations**:

> Text is always a nth-degree septum: it refers to a first-degree septum (like an image) or to other nth-degree septa (other interpretations

Can text never be a first-degree septum?

* What is a *source?*

Are we trying to build *the one notetaking app to rule them all?* If not, why do we have so many features? We should probably focus on a niche and cut down expected features to fit that niche.
