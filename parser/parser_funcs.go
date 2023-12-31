package parser

import (
	"fmt"

	"arc/lexer"
	"arc/log"
)

func (p *Parser) GetLexer() *lexer.Lexer {
	return p.lexer
}

func (p *Parser) toggleInfixFunc(t lexer.TokenType, enabled bool) {
	if enabled {
		if _, ok := p.disabledInfixParseFns[t]; !ok {
			return
		}
		p.infixParseFns[t] = p.disabledInfixParseFns[t]
		delete(p.disabledInfixParseFns, t)
	} else {
		if _, ok := p.infixParseFns[t]; !ok {
			return
		}
		p.disabledInfixParseFns[t] = p.infixParseFns[t]
		delete(p.infixParseFns, t)
	}
}
func (p *Parser) skipSemi() {
	if p.is(lexer.TokenSemicolon) {
		p.next()
	}
}
func (p *Parser) skipComma() {
	if p.is(lexer.TokenComma) {
		p.next()
	}
}

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
	if !p.is(t...) {
		p.error("expected %s, got %s instead", t, p.curr)
		return nil
	}
	p.next()

	return p.prev
}

//nolint:unused
func (p *Parser) expectPeek(t ...lexer.TokenType) {
	if !p.peekIs(t...) {
		p.error("expected next token to be one of %s, got %s instead", t, p.peek)
		return
	}
	p.next()
}

func (p *Parser) assertPrev(t ...lexer.TokenType) {
	if !p.prevIs(t...) {
		p.error("expected previous token to be one of %s, got %s instead", t, p.prev)
		return
	}
}

//nolint:unused
func (p *Parser) assert(t ...lexer.TokenType) {
	if !p.is(t...) {
		p.error("expected token to be one of %s, got %s instead", t, p.curr)
		return
	}
}

func (p *Parser) assertPeek(t ...lexer.TokenType) {
	if !p.peekIs(t...) {
		p.error("expected peek token to be one of %s, got %s instead", t, p.peek)
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

	log.Helper()
	err.Info = log.GetCallerInfo()

	panic(err)
}
func (p *Parser) error(fmt string, args ...any) {
	log.Helper()
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

func (p *Parser) unexpectedToken(curr *lexer.Token, reason ...string) {
	r := ""
	if len(reason) > 0 {
		r = reason[0]
		r = " - " + r
	}
	p.error(fmt.Sprintf("Unexpected token %s%s", curr, r))
}
