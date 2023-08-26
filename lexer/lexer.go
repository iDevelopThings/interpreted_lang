package lexer

import (
	"strings"
	"unicode"

	"arc/utilities"
)

type Lexer struct {
	input  string
	pos    *Position
	source string
	len    int
}

func NewLexer(input string) *Lexer {
	l := &Lexer{
		input: input,
		len:   len(input),
		pos: &Position{
			Line:   1,
			Column: 0,
			Abs:    0,
		},
	}

	return l
}

func (l *Lexer) newToken(tokenType TokenType, value string) *Token {
	tok := NewToken(value, tokenType)
	tok.Pos = NewTokenPosition(l.pos, len(value))

	return tok
}
func (l *Lexer) newTokenAndAdvance(tokenType TokenType, value string) *Token {
	tok := l.newToken(tokenType, value)
	l.advance()
	return tok
}

// Should only be used for testing.
func (l *Lexer) readAll() []*Token {
	t := utilities.NewTimer("Lexer.readAll()")
	defer t.StopAndLog()

	var tokens []*Token
	for tok := l.Next(); tok != nil; tok = l.Next() {
		if tok.Is(TokenEOF) {
			break
		}

		tokens = append(tokens, tok)

	}
	return tokens
}

func (l *Lexer) readNext() *Token {
	if l.pos.Abs >= len(l.input) {
		return l.newToken(TokenEOF, "")
	}

	l.consumeWhitespace()
	if l.pos.Abs >= len(l.input) {
		return l.newToken(TokenEOF, "")
	}

	ch := l.current()

	if ch == '/' {
		if l.peek() == '/' || l.peek() == '*' {
			l.readComment()
			return l.Next()
		}
	}

	// Prevent from running when matching `.0f` in a float.
	// if ch == '.' && !unicode.IsDigit(l.peek()) {
	// Attempt to match tokens using our mapping.
	// if tok, matched := l.matchToken(); matched {
	// 	return tok
	// }
	// }

	// If no mapped token is found, handle special cases like numbers and whitespace.
	switch {

	case isStringOpenChar(ch):
		return l.readString()

	case unicode.IsDigit(ch):
		// Peek to determine if it's a float
		if l.peek() == '.' && l.peek(2) != '.' {
			return l.readFloat()
		}
		return l.readInteger()

	case ch == '.' && unicode.IsDigit(l.peek()):
		// Direct float detection for numbers like .0f
		return l.readFloat()

	default:

		// First we'll try to match a symbol
		if tok := l.matchSymbols(); tok != nil {
			return tok
		}

		// Now we'll try and match some keywords
		// tok, matched := l.matchToken()
		// if matched {
		// 	return tok
		// }

		if ident := l.readIdentifier(); ident != nil {

			// We'll just do a quick check to see if this identifier is a keyword
			// if it is, we'll add the keyword type to the token

			for _, match := range keywordMatchTable {
				if ident.Value == match.Value {
					ident.AddType(match.Token)
				}
			}

			return ident
		}

		return l.newToken(TokenUnknown, string(ch))
	}
}

func (l *Lexer) Next() *Token {
	tok := l.readNext()
	if tok.Is(TokenUnknown) {
		panic("Unknown token: " + tok.Value)
	}
	return tok
}

func (l *Lexer) matchSymbols() *Token {
	// var symbolMatchTable = []tokenMatch{
	// 	{Value: ":", Token: TokenColon},
	// 	{Value: "::", Token: TokenColonColon},
	// }

	// We have our symbol match table like above(with more symbols ofc)
	// We need to try match a symbol in this table, but without consuming the input.
	// The table is re-ordered already so the longest symbols are first.

	var matchedSymbol = tokenMatch{"", TokenUnknown}

	for idx, symbol := range symbolMatchTable {
		if !strings.HasPrefix(l.remaining(), symbol.Value) {
			continue
		}

		matchedSymbol = symbolMatchTable[idx]
		break
	}

	if matchedSymbol.Value == "" {
		return nil
	}

	tok := l.newToken(matchedSymbol.Token, matchedSymbol.Value)

	// Advance the position by the length of the matched symbol
	for i := 0; i < len(matchedSymbol.Value); i++ {
		l.advance()
	}

	return tok
}

/*func (l *Lexer) matchToken() (*Token, bool) {
	maxLength := 0
	matchedToken := TokenType("")

	// Find the longest matching token
	for seq, tokenType := range tokenMap {
		// c := strings.ToLower(l.remaining())
		if len(seq) > maxLength && strings.HasPrefix(l.remaining(), seq) {
			maxLength = len(seq)
			matchedToken = tokenType
		}
	}

	if matchedToken != "" {
		start := l.pos.Abs
		for i := 0; i < maxLength; i++ {
			l.advance()
		}
		val := l.input[start : start+maxLength]
		// val = strings.ToLower(val)

		tok := l.newToken(matchedToken, val)

		// If we matched a keyword, we need to add the identifier type to the state also
		if _, ok := tokenKeywordMap[val]; ok {
			tok.AddType(TokenIdentifier)
		}
		return tok, true
	}

	return nil, false
}*/

func isIdentifierLetter(ch rune, afterFirst bool) bool {
	is := unicode.IsLetter(ch) || ch == '_'
	// we can have numbers after the first character
	if afterFirst {
		is = is || unicode.IsDigit(ch)
	}
	return is
}

// readIdentifier reads alphanumeric sequences starting with a letter.
func (l *Lexer) readIdentifier() *Token {
	start := l.pos.Abs

	read := 0
	for l.pos.Abs < l.len {
		if !isIdentifierLetter(l.current(), read > 0) {
			break
		}
		l.advance()
		read++
	}
	tok := l.newToken(TokenIdentifier, l.input[start:l.pos.Abs])

	return tok
}

// readInteger reads whole numbers.
func (l *Lexer) readInteger() *Token {
	start := l.pos.Abs
	for l.pos.Abs < l.len && unicode.IsDigit(l.current()) {
		l.advance()
	}
	return l.newToken(TokenInteger, l.input[start:l.pos.Abs])
}

// readFloat reads floating point numbers ending with 'f'.
func (l *Lexer) readFloat() *Token {
	start := l.pos.Abs
	if l.current() == '.' {
		l.advance()
	}
	for l.pos.Abs < l.len {
		if unicode.IsDigit(l.current()) {
			l.advance()
			continue
		}
		if l.current() == '.' {
			l.advance()
			continue
		}

		break
	}
	if l.current() != 'f' {
		return l.newToken(TokenUnknown, l.input[start:l.pos.Abs])
	}

	l.advance()

	return l.newToken(TokenFloat, l.input[start:l.pos.Abs])

	// if l.pos.Abs < l.len && l.current() == 'f' {
	// 	l.advance()
	// 	return l.newTokenAndAdvance(TokenFloat, l.input[start:l.pos.Abs])
	// }
	// return l.newTokenAndAdvance(TokenUnknown, l.input[start:l.pos.Abs])
}

func isStringOpenChar(ch rune) bool {
	return ch == '"' || ch == '\'' || ch == '`'
}

func (l *Lexer) readString() *Token {
	start := l.pos.Abs

	// We need to handle " ' and ` strings.
	strOpenKind := l.current()
	if !isStringOpenChar(strOpenKind) {
		return l.newToken(TokenUnknown, l.input[start:l.pos.Abs])
	}

	l.advance() // Consume the first quote
	for l.pos.Abs < l.len {
		if l.current() == strOpenKind {
			break
		}
		l.advance()
	}
	if l.current() != strOpenKind {
		return l.newToken(TokenUnknown, l.input[start:l.pos.Abs])
	}

	l.advance()

	tok := l.newToken(TokenString, l.input[start:l.pos.Abs])
	tok.Value = l.input[start+1 : l.pos.Abs-1]

	return tok
}

func (l *Lexer) readComment() {

	openType := string(l.current()) + string(l.peek())

	// Consume the first two characters
	l.advance()
	l.advance()

	if openType == "/*" && l.peek() == '*' {
		l.advance()
	}

	// Single line comment
	if openType == "//" {
		for l.pos.Abs < l.len && l.current() != '\n' {
			l.advance()
		}
		return
	}

	// Multi line comment
	for l.pos.Abs < l.len {
		if l.current() == '*' && l.peek() == '/' {
			l.advance()
			l.advance()
			return
		}
		l.advance()
	}

}

func (l *Lexer) consumeWhitespace() {
	if !unicode.IsSpace(l.current()) {
		return
	}
	// Any whitespace should just be consumed and not returned as a token.
	for unicode.IsSpace(l.current()) {
		l.advance()
	}
}

func (l *Lexer) debugDisplayTokens(w func(fmt string, args ...any), tokens []*Token) {
	const tabWidth = 4

	lines := strings.Split(l.input, "\n")

	for _, token := range tokens {
		line := lines[token.GetLine()-1] // since Line starts at 1
		col := token.GetColumn()

		// Replace tabs with spaces for alignment purposes
		adjustedLine := strings.ReplaceAll(line, "\t", strings.Repeat(" ", tabWidth))

		// Calculate the adjustment for the position marker
		beforeToken := line[:col]
		numTabsBeforeToken := strings.Count(beforeToken, "\t")
		adjustment := numTabsBeforeToken * (tabWidth - 1)

		prefix := strings.Repeat(" ", col+adjustment)

		marked := strings.Repeat("^", token.Pos.Length)

		w("%-3d: %s\n", token.GetLine(), adjustedLine)
		w("     %s%s  value=%s types=(%s)\n", prefix, marked, token.Value, token.Types)
		w("\n")
	}
}

func (l *Lexer) SetSource(path string) {
	l.source = path
}
func (l *Lexer) GetSource() string {
	return l.source
}
