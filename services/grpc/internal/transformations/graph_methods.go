package transformations

func AssembleDocForNode(node Node*) (doc string) {
	if (docIsPopulated) {
		return node.CheckpointDoc
	}

	// TODO
}