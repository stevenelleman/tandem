package transformations

import "web.microservice.shell/services/public-api/internal/models"

const (
	// NOTE: Transformations can be the same operation but interpretted differently in the UI
	//	eg Highlight and Take are the same operation but Highlight is surfaced on the doc while
	//	Take moves the piece to a separate doc
	REPLACE = iota
	DRAG
	HIGHLIGHT
	TAKE
	EMBED
)

func Call(parent string, args *models.FxnArgs) (string, *models.FxnArgs) {
	var child, iregion string
	var istart1, iend1, istart2 int

	// NOTE: How to make functions extensible and discoverable?
	// Idea: Make a scope of available transformations. Transformations have a special type.
	// Text methods can be used on them. How to define types?
	switch args.FxnId {
	case REPLACE:
		child, iregion,  istart1, iend1 = Replace(parent, args.Source, args.OffsetStart1, args.OffsetEnd1)
	case DRAG:
		child, istart1, iend1, istart2 = Drag(parent, args.OffsetStart1, args.OffsetEnd1, args.OffsetStart2)
	case HIGHLIGHT:
		child = Highlight(parent, args.OffsetStart1, args.OffsetEnd1)
	case TAKE:
		child = Take(parent, args.OffsetStart1, args.OffsetEnd1)
	case EMBED:
		child, iregion, istart1, iend1 = Embed(parent, args.Source, args.OffsetStart1, args.OffsetEnd1)
	}

	iargs := &models.FxnArgs{
		FxnId:        args.FxnId,
		Source:       iregion,
		OffsetStart1: istart1,
		OffsetEnd1:   iend1,
		OffsetStart2: istart2,
	}

	return child, iargs
}

// Take in some source node, create an edge and a target node
func Transform(parent *models.Node, args *models.FxnArgs) (*models.Node, *models.Edge) {
	targetDoc, iargs := Call(parent.Document, args)
	child := models.ConstructNode(targetDoc, parent)
	edge := models.ConstructEdge(args, iargs, parent.Id, child.Id)
	child.EdgeIn = edge.Id
	return child, edge
}
