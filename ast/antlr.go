package ast

type TokenRange struct {
	StartLine int
	StartCol  int

	StopLine int
	StopCol  int

	// The left bound column of the range
	RangeStartCol int
	// The right bound column of the range
	RangeStopCol int
}

func (self *AstNode) GetTokenRange() *TokenRange {
	ctx := self.GetToken()
	if ctx == nil {
		return nil
	}

	start := ctx.GetStart()
	stop := ctx.GetEnd()

	r := &TokenRange{
		StartLine: start.GetLine(),
		StartCol:  start.GetColumn(),
		StopLine:  stop.GetLine(),
		StopCol:   stop.GetColumn(),

		RangeStartCol: 0,
		RangeStopCol:  0,
	}

	// if we have this input:
	// (u User)
	//  ^ ^
	// start/stop use the character at which the token starts, for the column
	// regardless of their names
	// So now we have to find the correct start/stop column offsets

	// Essentially mapping it to
	// u User
	// ^    ^
	r.RangeStartCol = r.StartCol
	r.RangeStopCol = r.StopCol + ctx.Pos.Length

	return r
}

func (self *TokenRange) ZeroIndexed() *TokenRange {
	self.StartLine = max(0, self.StartLine-1)
	// self.StartCol = max(0, self.StartCol-1)

	self.StopLine = max(0, self.StopLine-1)
	// self.StopCol = max(0, self.StopCol-1)

	// self.RangeStartCol = max(0, self.RangeStartCol-1)
	// self.RangeStopCol = max(0, self.RangeStopCol-1)

	return self
}
