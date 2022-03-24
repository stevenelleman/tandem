ArchGraph Proposal: A Direction for SourceGraph:

A sourcegraph exists on two orthogonal dimensions: time and space.

## Time:

* Works like git. Diffs are saved, not documents. The entire sourcegraph operates on the same version control, so if you make a change to a single document, that will apply as a diff to the entire SourceGraph.
  *  In practice, I suspect changes to single documents will be very rare anyway, for reasons we will discuss.
  * Later, you might be able to split versions of the SourceGraph and merge them similar to Git. Automatic merging and managing collisions might be hard, though. This will help especially with multi-user Sourcegraphs.
  * Later, monetization might depend on the system of diffs by keeping track of contributers and who contributed what, when, to what version of the original branch.
  * Also, we should define diffs as transformations and allow users to customize transformations (talk to Steven about this!)
  * For now, though, we can just focus on allowing diffs and branching. 



## Space:

    * There are two labels of things: nodes and arrows.

### Node:
    * A node can do just four things:
     1. Have a title
     2. Have content (just unformatted text to start with)
     3. Be labelled (have any number of labels applied to them)
     4. Have relationships to other nodes, represented by arrows

     A node **must** have a title, but it doesn't have to do any of the other things. A node with no content is empty. A node with no label is generic. A node with no relationships is singular. A node that is empty, generic, and singular is called a **nub**. Nubs are important because SourceGraph generates them automatically. We'll get to that.

     Remember that nubs still have titles! But they don't have contents, labels, or relationships to other nodes.

### Arrow:
     * An arrow can do just two things:
     1. Point from one node to another node
     2. Be labeld

     The second concept  is huge. **Relationships can be labeld, because there are many different kinds of relationships between concepts.**

## Labels

     A **label** is a label you apply to a node. This represents an "is-a" relationship.  For example, say you were comparing car models, so you add a node called "Ford Escape" and apply the label "car" to it. This makes sense because a Ford Escape **is a** car.

     Labels can also be applied to relationships. Say you add the node **Ford** with the label **car manufacturer** you could draw arrows leading from Ford to Ford Escape, Ford Focus, Ford Fiest, and others. Each of these arrows would be tagged with the **made-by** tag to represent the label of relationships that arrow indicates.

     The most important fact about labels is that they are actually nodes. Whenever you add a new label to an arrow or node, SourceGraph creates a nub with the title of that label. In the previous example, SourceGraph would create a car manufacturer nub, a made-by nub, and a car  nub. If you want, you can turn these nubs into nodes, by adding content, applying labels to them, or adding relationships between them and other nodes.

## Example:

     Say you want to document the TypeScript programming language. You make a node called "TypeScript" with the label "Programming Language." Because labels are nodes, SourceGraph creates a nub with the title "Programming Language" that you can later modify.

     Next, you add a node called "Static Typing" and connect it with Typescript with an arrow. You label that arrow "is a feature of," and SourceGraph creates a nub entitled "is a feature of." You keep doing this with various features, updates, IDEs, employees, and you suddenly get a rich representation not only of the programming language itself, but also how it implements ideas from computer science itself, ideas like compilers, updates, and features. This higher level of abstraction is generated naturally by you when you label nodes and arrows, because each label becomes its own node. 

     See a poor-quality image of this example on paper! https://imgur.com/a/UBBsdOQ
     
## Example II:

 This example demonstrates how this system can be used to apply many different _ways of seeing something_ to a single subject. Imagine you are charting the history of science. You can create a node for an inventor, say "Euclid", and label it with the "person" label. Next, you create nodes labelled "invention" and tie them to Euclid with an arrow labelled "created." You create nodes and relationships in a like manner for everything Euclid did, all those he influenced and those who influenced him, and even his time period "belonged to --> 4th c. BC." 
 
 Now, you research other figures in the History of Science in a similar way. You can now filter your graph in at least the following ways:
 
 * By Time period: search for all nodes that point toward the "15th century" node with a "belonged to" arrow.
 * 



# Questions:

     * Right now it looks like arrows only connect nodes to nodes, which is different from the concept of links (which connect specific pieces of text to pages.) How do you solve this? This might be okay if nodes are atomic, but it could quickly get annoying.
        * What problems do links solve? Well, I think they solve the same problem that tagging arrows tries to solve: they contextualize the link in relation to the node, they show not only **that** the documents relate (which is what an arrow in a mathematical graph does), but also **how** they relate. So I think a labelled, directional arrow would be sufficient for this, especially if the nodes were atomic, but perhaps as an *extra feature* you could make links tie to text in either the sending or receiving node (with the receiving node being the node the arrow points to, and the sending node being the node that the arrow points from).
        * How would it feel to add to other SourceGraphs?
        * How would it feel to write a novel with this?
        * Does this generalize to many usecases? Is there something the user is missing here?

# Ideas:

        * Use basic ML to have SourceGraph automatically suggest tags for documents. Let's save the user as much work as possible!



# Pros:

        * Elegance. Because labels are nodes, everything is expressed by the fewest labels of things possible with the fewest features possible.
        * Emergent hierarchy. By using sourcegraph to document something specific, you naturally generate a hierarchy for the more general version of that thing which you can use to understand more examples of speific things in the future.
        * Easy to search. Because arrows and nodes can have as many labels as the user wants, you can easily implement powerful boolean search on both nodes and arrows.
        * Extensible to any field. Because labels are generic, this system can represent any things and any relationships the user wants it to. It is a highly-general knowledge management tool.

# Cons:

        * Nubs generated by arrow labels are clunky.
        * Forces the user to develop a philosophical hierarchy when they might not need to.
        * Might be hard to support top-down hierarchies.
        * Does not represent "insideness" very well. Imagine the "biology" subject has 20 subsubjects. You need to drag an "inside" arrow from all subsubjuects to Biology, which would be clunky if done manually. "Select All" and other like tools could make the process faster, but this still feels inelegant.
