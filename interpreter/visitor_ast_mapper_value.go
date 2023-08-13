package interpreter

import (
	"strconv"

	"github.com/antlr4-go/antlr/v4"

	"interpreted_lang/ast"
	"interpreted_lang/grammar"
)

func (self *AstMapper) VisitDictFieldKey(ctx *grammar.DictFieldKeyContext) interface{} {
	if v := ctx.ID(); v != nil {
		return ast.NewIdentifier(ctx)
	} else if v := ctx.String_(); v != nil {
		str := self.VisitAnyValue(v).(*ast.Literal)
		return ast.NewIdentifierWithValue(ctx, str.Value.(string))
	} else {
		panic("unknown dict field key")
	}
}

func (self *AstMapper) VisitDict(ctx *grammar.DictContext) interface{} {
	dict := &ast.DictionaryInstantiation{
		AstNode: ast.NewAstNode(ctx),
		Fields:  map[string]ast.Expr{},
	}

	for _, field := range ctx.AllDictFieldAssignment() {
		assignment := self.Visit(field.GetVal())
		key := self.Visit(field.GetKey()).(*ast.Identifier)
		dict.Fields[key.Name] = assignment.(ast.Expr)
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
		return &ast.Literal{
			ast.NewAstNode(ctx),
			ast.LiteralKindString,
			ctx.GetText()[1 : len(ctx.GetText())-1],
		}

	case grammar.SimpleLangLexerVALUE_BOOL:
		return &ast.Literal{
			ast.NewAstNode(ctx),
			ast.LiteralKindBoolean,
			ctx.GetText() == "true",
		}

	case grammar.SimpleLangLexerVALUE_INTEGER:
		integer, err := strconv.Atoi(ctx.GetText())
		if err != nil {
			panic(err)
		}
		return &ast.Literal{
			ast.NewAstNode(ctx),
			ast.LiteralKindInteger,
			integer,
		}

	case grammar.SimpleLangLexerVALUE_FLOAT:
		float, err := strconv.ParseFloat(ctx.GetText()[0:len(ctx.GetText())-1], 64)
		if err != nil {
			panic(err)
		}
		return &ast.Literal{
			ast.NewAstNode(ctx),
			ast.LiteralKindFloat,
			float,
		}

	case grammar.SimpleLangLexerVALUE_NULL:
		return &ast.Literal{
			ast.NewAstNode(ctx),
			ast.LiteralKindNull,
			nil,
		}

	case grammar.SimpleLangParserID:
		return &ast.VarReference{
			ast.NewAstNode(ctx),
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
	arr := &ast.ArrayInstantiation{
		AstNode: ast.NewAstNode(ctx),
		Values:  make([]ast.Expr, 0),
	}

	for _, value := range ctx.AllListElement() {
		listElement := self.Visit(value.GetVal())

		arr.Values = append(arr.Values, listElement.(ast.Expr))
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
