package transformations

import (
	"math/rand"
)

func GenerateUUID() (id int) {
	return utils.UUIDv4().String()
}

func ConstructNode(doc string, parent *Node) (*Node) {
	return &Node {
		Id: GenerateUUID(),
		Document: doc,
		Scopes: parent.Scopes, // Retain the same scopes as parent

		CreatedAt: time.Now(), // TODO: Express as a PreInsert so db clock is source of truth
	}
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

func ConstructScope(node *Node) (*Node) {
	// TODO
}

// Create the first node in a document, with no source edges
func CreateNewDocument() (*Node) {
	node := &Node {
		Id: GenerateUUID(),
		CreatedAt: time.Now(), // TODO: Express as a PreInsert so db clock is source of truth
	}

	// Create a new scope for the new document
	scope := &Scope {
		ScopeId: GenerateUUID(),
		Nodes: [node.NodeId]
	}

	node.scopes = append(node.scopes, scope)
	return node
}

// Take in some source node, create an edge and a target node
func Transform(parent *models.Node, args *models.FxnArgs) (*models.Node, *models.Edge) {
	targetDoc, iargs := Call(parent.Document, args)
	child := models.ConstructNode(targetDoc, parent)
	edge := models.ConstructEdge(args, iargs, parent.Id, child.Id)
	child.EdgeIn = edge.Id
	return child, edge
}

func ExtractDocFromTransformation(edge Edge*) (doc string) {
	var region string := "" // Do we need this?
	var args Args* := edge.Args
	var sourceNode Node* = GetNodeById(graph, edge.SourceNodeId) // TODO set up preprocessing to access graph as go object, cache as necessary
	var sourceDoc string := AssembleDocForNode(sourceNode)

	targetDoc, inverseDoc, invArgs := Call(region, sourceDoc, args)
	return targetDoc
}
func AssembleDocForNode(node Node*) (doc string) {
	if (node.DocIsPopulated) {
		return node.CheckpointDoc
	}

	// Extract doc from source edge. Recursively accumulate it.
	return ExtractDocFromTransformation(currentNode.EdgesIn[0]) // Currently assumes one parent! Will be a challenge to deal with multiple
}
