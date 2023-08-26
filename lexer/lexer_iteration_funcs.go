package lexer

import (
	"unicode/utf8"
)

func (l *Lexer) peek(ni ...int) rune {
	n := 1
	if len(ni) > 0 {
		n = ni[0]
	}
	if l.pos.Abs+n >= len(l.input) {
		return 0 // Null character
	}
	return rune(l.input[l.pos.Abs+n])
}

func (l *Lexer) PeekNext() *Token {
	tokens := l.PeekNextN(1)
	if len(tokens) == 0 {
		return nil
	}
	return tokens[0]
}

func (l *Lexer) PeekNextN(n int) []*Token {
	startPos := *l.pos

	var tokens []*Token

	for i := 0; i < n; i++ {
		tok := l.Next()
		tokens = append(tokens, tok)
	}

	*l.pos = startPos

	return tokens
}

func (l *Lexer) current() rune {
	if l.pos.Abs >= len(l.input) {
		return 0 // Null character
	}
	return rune(l.input[l.pos.Abs])
}

func (l *Lexer) remaining() string {
	return l.input[l.pos.Abs:]
}

func (l *Lexer) advance() {
	_, size := utf8.DecodeRuneInString(string(l.current()))

	if l.current() == '\n' {
		l.pos.Line++
		// -1 because advance() will increment it to 0
		l.pos.Column = -1
	}

	l.pos.Abs += size
	l.pos.Column += size
}
