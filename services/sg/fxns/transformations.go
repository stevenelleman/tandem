package fxns

const (
	REPLACE = iota
	TRANSLCUDE
	DRAG
	HIGHLIGHT
)

// Each function should conform to this interface
// Each function should have a constructor of some sort _outside_ of the interface with specific arguments
type Transformation interface {
	Call(string, string, *FxnArgs) (string, string, *FxnArgs)
}

type FxnArgs struct {
	FxnId        int
	OffsetStart1 int
	OffsetEnd1   int
	OffsetStart2 int
}

// Return the result and inverse
func Call(region0, src0 string, args *FxnArgs) (region1, iregion string, iargs *FxnArgs) {
	var istart1, iend1, istart2 int

	// TODO: How to easily extend this function? Wish there was a way to do a function lookup
	switch args.FxnId {
	case REPLACE:
		region1, iregion, istart1, iend1 = ReplaceInverse(region0, src0, args.OffsetStart1, args.OffsetEnd1)
	case TRANSLCUDE:
		region1, iregion, istart1, iend1 = TranscludeInverse(region0, src0, args.OffsetStart1, args.OffsetEnd1)
	case DRAG:
		region1, istart1, iend1, istart2 = DragInverse(region0, args.OffsetStart1, args.OffsetEnd1, args.OffsetStart2)
	case HIGHLIGHT:
		region1 = Highlight(region0, args.OffsetStart1, args.OffsetEnd1)
	}

	iargs = &FxnArgs{
		FxnId:        args.FxnId,
		OffsetStart1: istart1,
		OffsetEnd1:   iend1,
		OffsetStart2: istart2,
	}

	return region1, iregion, iargs
}
