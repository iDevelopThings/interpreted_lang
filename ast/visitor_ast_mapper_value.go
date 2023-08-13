package ast

import (
	"strconv"

	"github.com/antlr4-go/antlr/v4"

	"interpreted_lang/grammar"
)

func (self *AstMapper) VisitDictFieldKey(ctx *grammar.DictFieldKeyContext) interface{} {
	if v := ctx.ID(); v != nil {
		return NewIdentifier(ctx)
	} else if v := ctx.String_(); v != nil {
		str := self.VisitAnyValue(v).(*Literal)
		return NewIdentifierWithValue(ctx, str.Value.(string))
	} else {
		panic("unknown dict field key")
	}
}

func (self *AstMapper) VisitDict(ctx *grammar.DictContext) interface{} {
	dict := &DictionaryInstantiation{
		AstNode: NewAstNode(ctx),
		Fields:  map[string]Expr{},
	}

	for _, field := range ctx.AllDictFieldAssignment() {
		assignment := self.Visit(field.GetVal())
		key := self.Visit(field.GetKey()).(*Identifier)
		dict.Fields[key.Name] = assignment.(Expr)
	}

	return dict
}

func (self *AstMapper) VisitString(ctx *grammar.StringContext) interface{} {
	return self.VisitAnyValue(ctx)
}

func (self *AstMapper) VisitInt(ctx *grammar.IntContext) interface{} {
	return self.VisitAnyValue(ctx)
}

func (self *AstMapper) VisitFloat(ctx *grammar.FloatContext) interface{} {
	return self.VisitAnyValue(ctx)
}

func (self *AstMapper) VisitBool(ctx *grammar.BoolContext) interface{} {
	return self.VisitAnyValue(ctx)
}

func (self *AstMapper) VisitNull(ctx *grammar.NullContext) interface{} {
	return self.VisitAnyValue(ctx)
}

func (self *AstMapper) VisitAnyValue(ctx antlr.ParserRuleContext) any {

	if ctx == nil {
		return nil
	}

	if v, ok := ctx.(*grammar.DictContext); ok {
		return self.VisitDict(v)
	}

	switch ctx.GetStart().GetTokenType() {
	case grammar.SimpleLangLexerDOUBLE_QUOUTE_STRING,
		grammar.SimpleLangLexerSINGLE_QUOUTE_STRING,
		grammar.SimpleLangLexerBACKTICK_STRING:
		return &Literal{
			NewAstNode(ctx),
			LiteralKindString,
			ctx.GetText()[1 : len(ctx.GetText())-1],
		}

	case grammar.SimpleLangLexerVALUE_BOOL:
		return &Literal{
			NewAstNode(ctx),
			LiteralKindBoolean,
			ctx.GetText() == "true",
		}

	case grammar.SimpleLangLexerVALUE_INTEGER:
		integer, err := strconv.Atoi(ctx.GetText())
		if err != nil {
			panic(err)
		}
		return &Literal{
			NewAstNode(ctx),
			LiteralKindInteger,
			integer,
		}

	case grammar.SimpleLangLexerVALUE_FLOAT:
		float, err := strconv.ParseFloat(ctx.GetText()[0:len(ctx.GetText())-1], 64)
		if err != nil {
			panic(err)
		}
		return &Literal{
			NewAstNode(ctx),
			LiteralKindFloat,
			float,
		}

	case grammar.SimpleLangLexerVALUE_NULL:
		return &Literal{
			NewAstNode(ctx),
			LiteralKindNull,
			nil,
		}

	case grammar.SimpleLangParserID:
		return &VarReference{
			NewAstNode(ctx),
			ctx.GetText(),
		}

	default:
		tokType := ctx.GetStart().GetTokenType()
		panic("Unknown value, tokenType: " + strconv.Itoa(tokType) + " value string: \n" + ctx.GetText())
	}
}

func (self *AstMapper) VisitValuePrimary(ctx *grammar.ValuePrimaryContext) interface{} {
	if value := ctx.Value(); value != nil {
		return self.VisitValue(value.(*grammar.ValueContext))
	}

	return self.VisitAnyValue(ctx)
}

func (self *AstMapper) VisitList(ctx *grammar.ListContext) interface{} {
	arr := &ArrayInstantiation{
		AstNode: NewAstNode(ctx),
		Values:  make([]Expr, 0),
	}

	for _, value := range ctx.AllListElement() {
		listElement := self.Visit(value.GetVal())

		arr.Values = append(arr.Values, listElement.(Expr))
	}

	return arr
}

func (self *AstMapper) VisitValue(ctx *grammar.ValueContext) interface{} {
	if dict := ctx.Dict(); dict != nil {
		return self.VisitDict(dict.(*grammar.DictContext))
	}
	if list := ctx.List(); list != nil {
		return self.VisitList(list.(*grammar.ListContext))
	}
	if obj := ctx.ObjectInstantiation(); obj != nil {
		return self.VisitObjectInstantiation(obj.(*grammar.ObjectInstantiationContext))
	}

	return self.VisitAnyValue(ctx)
}
