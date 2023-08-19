package parser

import (
	"arc/lexer"
	"arc/log"
)

func (p *Parser) next() {
	if p.curr != nil {
		p.prev = p.curr
	}
	if p.peek != nil {
		p.curr = p.peek
	}
	p.peek = p.lexer.Next()
}

func (p *Parser) is(t ...lexer.TokenType) bool {
	if p.curr == nil {
		return false
	}

	for _, tokenType := range t {
		if p.curr.Is(tokenType) {
			return true
		}
	}

	return false
}

func (p *Parser) peekIs(t ...lexer.TokenType) bool {
	if p.peek == nil {
		return false
	}

	for _, tokenType := range t {
		if p.peek.Is(tokenType) {
			return true
		}
	}

	return false
}

func (p *Parser) prevIs(t ...lexer.TokenType) bool {
	if p.prev == nil {
		return false
	}

	for _, tokenType := range t {
		if p.prev.Is(tokenType) {
			return true
		}
	}

	return false
}

func (p *Parser) expect(t ...lexer.TokenType) *lexer.Token {
	// log.Log.MarkAsHelperFunc(1)
	if !p.is(t...) {
		p.error("expected %s, got %s instead", t, p.curr)
		return nil
	}
	p.next()

	return p.prev
}

func (p *Parser) expectPeek(t ...lexer.TokenType) {
	// log.Log.MarkAsHelperFunc(1)
	if !p.peekIs(t...) {
		p.error("expected next token to be one of %s, got %s instead", t, p.peek)
		return
	}
	p.next()
}

func (p *Parser) assertPrev(t ...lexer.TokenType) {
	// log.Log.MarkAsHelperFunc(1)
	if !p.prevIs(t...) {
		p.error("expected previous token to be one of %s, got %s instead", t, p.prev)
		return
	}
}

func (p *Parser) tokenError(token *lexer.Token, fmt string, args ...any) {
	err := &ParserError{
		Fmt:    fmt,
		Args:   args,
		Token:  token,
		Parser: p,
	}

	log.Log.MarkAsHelperFunc(1)
	err.Info = log.Log.CallerInfo(1)

	panic(err)
}
func (p *Parser) error(fmt string, args ...any) {
	log.Log.MarkAsHelperFunc(1)
	p.tokenError(p.curr, fmt, args...)
}

func (p *Parser) safeLoop(predicate func() bool, cb func()) {

	i := 0
	maxIteration := 100

	for predicate() {
		cb()
		i++
		if i > maxIteration {
			p.error("infinite loop detected")
			return
		}
	}
}
