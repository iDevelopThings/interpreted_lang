package lexer

type LexingError struct {
	Message string
	Pos     *TokenPosition
}

func (e LexingError) GetPosition() any {
	return e.Pos
}

func (e LexingError) GetMessage() string {
	return e.Message
}

func (e LexingError) Error() string {
	return e.Message
}

// func (e LexingError) GetRelatedNode() NodeLike {
// 	return e.Pos
// }
