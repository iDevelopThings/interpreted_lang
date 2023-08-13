package grammar

import (
	"github.com/antlr4-go/antlr/v4"
)

type SimpleLangLexerBase struct {
	*antlr.BaseLexer

	expectingObjectKeyword bool
}

func (p SimpleLangLexerBase) isType() bool {
	return false
}

type SimpleLangParserBase struct {
	*antlr.BaseParser
}

func (p *SimpleLangParserBase) closingBracket() bool {
	return false
	// stream := p.GetTokenStream()
	// prevTokenType := stream.LA(1)
	// return prevTokenType == GoParserR_PAREN || prevTokenType == GoParserR_CURLY;
}
func (p *SimpleLangParserBase) isHttpKeyword() bool {
	return false
}

func (p SimpleLangParserBase) isType() bool {
	return false
}

func (p SimpleLangParserBase) isObjectKeyword() bool {
	word := p.GetCurrentToken().GetText()

	keywords := []string{"body", "string"}
	for _, keyword := range keywords {
		if word == keyword {
			return true
		}
	}
	return false
}
