package transformations

const (
	REPLACE = iota
	TRANSCLUDE
	DRAG
	HIGHLIGHT
)

type FxnArgs struct {
	FxnId        int
	OffsetStart1 int
	OffsetEnd1   int
	OffsetStart2 int
}

// NOTE: How to get all functions to conform to signature: Call(string, string, *FxnArgs) (string, string, *FxnArgs) ?
func Call(region0, src0 string, args *FxnArgs) (string, string, *FxnArgs) {
	var region1, iregion string
	var istart1, iend1, istart2 int

	// NOTE: How to make functions extensible and discoverable?
	switch args.FxnId {
	case REPLACE:
		region1, iregion, istart1, iend1 = Replace(region0, src0, args.OffsetStart1, args.OffsetEnd1)
	case TRANSCLUDE:
		region1, iregion, istart1, iend1 = Transclude(region0, src0, args.OffsetStart1, args.OffsetEnd1)
	case DRAG:
		region1, istart1, iend1, istart2 = Drag(region0, args.OffsetStart1, args.OffsetEnd1, args.OffsetStart2)
	case HIGHLIGHT:
		region1 = Highlight(region0, args.OffsetStart1, args.OffsetEnd1)
	}

	iargs := &FxnArgs{
		FxnId:        args.FxnId,
		OffsetStart1: istart1,
		OffsetEnd1:   iend1,
		OffsetStart2: istart2,
	}

	return region1, iregion, iargs
}
