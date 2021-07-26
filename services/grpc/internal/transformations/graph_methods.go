package transformations

import (
	"math/rand"
)

func GenerateUUID() (id int) {
	return rand.Int()
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
