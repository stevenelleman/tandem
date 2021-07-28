package graph

import (
	"math/rand"
)

// ----------------------------------------------------
// Methods for user functionality

// Create the first node in a document, with no source edges
func CreateNewDocument() (*Node) {
	node := ConstructEdge(false, "", nil)
	// Create a new scope for the new document
	scope := ConstructScope([node])
	node.scopes = append(node.scopes, scope)
	return node
}

// Take in some source node, create an edge and a target node
func Transform(parent *models.Node, args *models.FxnArgs) (*models.Node, *models.Edge) {
	targetDoc, iargs := Call(parent.Document, args)
	child := models.ConstructNode(true, targetDoc, parent)
	edge := models.ConstructEdge(args, iargs, parent.Id, child.Id)
	child.EdgeIn = edge.Id
	return child, edge
}

// Provide the document for any existing node
func AccumulateDocForNode(node Node*) (doc string) {
	if (node.DocIsPopulated) {
		return node.CheckpointDoc
	}

	// Accumulate doc from source edge. Recursively accumulate it.
	return AccumulateDocForEdge(currentNode.EdgesIn[0]) // Currently assumes one parent! Will be a challenge to deal with multiple
}

// Create a new scope from an existing node
func ForkDocument(node Node*) () {
	scope := ConstructScope([node])
	node.scopes := append(node.scopes, scope)
	return
}

// ----------------------------------------------------
// Internal Methods

func GenerateUUID() (id int) {
	return utils.UUIDv4().String()
}

func ConstructNode(populateDoc bool, doc string, parent *Node) (*Node) {
	node := &Node {
		Id: GenerateUUID(),
		DocIsPopulated: false
		CreatedAt: time.Now(), // TODO: Express as a PreInsert so db clock is source of truth
	}
	if (populateDoc) {
		node.Document = doc
		node.DocIsPopulated = true
	}
	if (parent) {
		node.scopes = parent.scopes // Retain the same scopes as parent
	}
	return node
}

func ConstructEdge(args, iargs *FxnArgs, parentId, childId string) (*Edge) {
	return &Edge{
		Id: GenerateUUID(),

		Args: args,
		InverseArgs: iargs,

		ParentNodeId: parentId,
		ChildNodeId: childId,

		CreatedAt: time.Now(), // TODO: Express as a PreInsert so db clock is source of truth
	}
}

func ConstructScope(nodes []*Node) (*Node) {
	return &Scope {
		Id: GenerateUUID(),
		Nodes: nodes,
		CreatedAt: time.Now(), // TODO: Express as a PreInsert so db clock is source of truth
	}
}

func AccumulateDocForEdge(edge Edge*) (doc string) {
	region := "" // Do we need this?
	args := edge.Args
	sourceNode = GetNodeById(graph, edge.SourceNodeId) // TODO set up preprocessing to access graph as go object, cache as necessary
	sourceDoc := AccumulateDocForNode(sourceNode)

	targetDoc, inverseDoc, invArgs := Call(region, sourceDoc, args)
	return targetDoc
}
