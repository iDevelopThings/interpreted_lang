package lexer

import (
	"strconv"
)

type Position struct {
	Line   int
	Column int
	Abs    int
}

func (p Position) String() string {
	return "line=" + strconv.Itoa(p.Line) + " - column=" + strconv.Itoa(p.Column) + " - abs=" + strconv.Itoa(p.Abs)
}

func (p *Position) GetLine() int   { return p.Line }
func (p *Position) GetColumn() int { return p.Column }
func (p *Position) GetAbs() int    { return p.Abs }

type TokenPosition struct {
	Start  *Position
	End    *Position
	Length int
}

// TODO: We should remove the pointers here, we could get some gains
// Once the objects are assigned to the token, they never ever change
func NewTokenPosition(end *Position, len int) *TokenPosition {

	tp := &TokenPosition{
		Start: &Position{
			Line:   end.Line,
			Column: end.Column - len,
			Abs:    end.Abs - len,
		},
		End: &Position{
			Line:   end.Line,
			Column: end.Column,
			Abs:    end.Abs,
		},
		Length: len,
	}

	return tp
}

func (p *TokenPosition) String() string {
	return "start=" + p.Start.String() + " - end=" + p.End.String() + " - len=" + strconv.Itoa(p.Length)
}

func (p *TokenPosition) GetStart() *Position { return p.Start }
func (p *TokenPosition) GetEnd() *Position   { return p.End }
func (p *TokenPosition) GetLength() int      { return p.Length }
