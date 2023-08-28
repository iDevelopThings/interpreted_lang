package interpreter

import (
	"os"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"testing"

	"arc/ast"
	"arc/utilities"
)

type TestInterpreterExpressionsTestSuite struct {
	suite.Suite
}

func TestInterpreterExpressionsSuite(t *testing.T) {
	suite.Run(t, new(TestInterpreterExpressionsTestSuite))
}

func (suite *TestInterpreterExpressionsTestSuite) SetupTest() {
	// Set an environment variable to indicate that we are running tests

	os.Setenv("INTERPRETER_TEST", "true")

}

func (suite *TestInterpreterExpressionsTestSuite) TearDownSuite() {

}

func (suite *TestInterpreterExpressionsTestSuite) Test_BasicExpressions() {

	tests := []struct {
		input    string
		rt       string
		expected any
	}{
		{rt: "int", input: "2+value(2)", expected: 4},
		{rt: "int", input: "2+2", expected: 4},
		{rt: "int", input: "2*2", expected: 4},
		{rt: "int", input: "2-2", expected: 0},
		{rt: "int", input: "2/2", expected: 1},
		{rt: "int", input: "(2*3)+4", expected: 10},
		{rt: "int", input: "2*3+4", expected: 10},
		{rt: "int", input: "2+3*4", expected: 14},
		{rt: "int", input: "2*(1+2)", expected: 6},
		{rt: "int", input: "(2+3)*4", expected: 20},
		{rt: "int", input: "2+(3*4)", expected: 14},
		{rt: "int", input: "2*3-4", expected: 2},
		{rt: "int", input: "2-3*4", expected: -10},
		{rt: "int", input: "(2-3)*4", expected: -4},
		{rt: "int", input: "2-(3*4)", expected: -10},
		{rt: "float", input: "2*3/4.0f", expected: 1.5},
		{rt: "float", input: "2.0f/3.0f*4.0f", expected: 2.6666666666666665},
		{rt: "int", input: "2/3*4", expected: 0},
		{rt: "float", input: "(2*3)/4.0f", expected: 1.5},
		{rt: "float", input: "2*(3.0f/4.0f)", expected: 1.5},
		{rt: "int", input: "10/3*2+4-(6-2)*(3+1)/2", expected: 2},
		{rt: "int", input: "(4+5*(2-3))/2-(6+4*2)/2", expected: -7},
	}

	for _, test := range tests {
		suite.Run(test.input, func() {

			engine := NewTestingInterpreterEngine()
			engine.DisableTypeChecker = true
			engine.LoadSourceFromString(`
func value(x ` + test.rt + `) ` + test.rt + ` { return x } 
func eval() ` + test.rt + ` { return ` + test.input + ` }
`)
			engine.ProcessScripts()

			script := engine.SourceFiles[0]
			evalFunc := script.GetFunc("eval")

			w := utilities.NewIndentWriter(os.Stdout, " ")
			evalFunc.PrintTree(w.(*utilities.IndentWriter))

			assert.NotNil(suite.T(), evalFunc)

			call := &ast.CallExpression{
				AstNode:  ast.NewAstNode(nil),
				Function: ast.NewIdentifierWithValue(nil, "eval"),
			}

			result := engine.Evaluator.MustEvalValue(call)

			assert.IsType(suite.T(), new(ast.RuntimeValue), result)

			rtResult := result.(*ast.RuntimeValue)
			assert.Equal(suite.T(), test.rt, rtResult.TypeName)

			switch test.expected.(type) {
			case int:
				assert.Equal(suite.T(), test.expected, rtResult.Value)
			case float64:
				assert.Equal(suite.T(), test.expected, rtResult.Value)
			}

		})
	}

}

func (suite *TestInterpreterExpressionsTestSuite) Test_Assignment() {

	tests := []struct {
		input    string
		rt       string
		expected any
	}{
		{rt: "int", input: "var x int = 0\nx++", expected: 1},
		{rt: "int", input: "var x int = 0\nx += 10", expected: 10},
		{rt: "int", input: "var x int = 0\nx -= 10", expected: -10},
		{rt: "int", input: "var x int = 2\nx *= 2", expected: 4},
		{rt: "int", input: "var x int = 4\nx /= 2", expected: 2},
		{rt: "int", input: "var x int = 2.0f", expected: 2},
	}

	for _, test := range tests {
		suite.Run(test.input, func() {

			engine := NewTestingInterpreterEngine()
			engine.DisableTypeChecker = true
			engine.LoadSourceFromString(`
func eval() ` + test.rt + ` {
	` + test.input + `
	return x 
}
`)
			engine.ProcessScripts()

			script := engine.SourceFiles[0]
			evalFunc := script.GetFunc("eval")

			w := utilities.NewIndentWriter(os.Stdout, " ")
			evalFunc.PrintTree(w.(*utilities.IndentWriter))

			assert.NotNil(suite.T(), evalFunc)

			call := &ast.CallExpression{
				AstNode:  ast.NewAstNode(nil),
				Function: ast.NewIdentifierWithValue(nil, "eval"),
			}

			result := engine.Evaluator.MustEvalValue(call)

			assert.IsType(suite.T(), new(ast.RuntimeValue), result)

			rtResult := result.(*ast.RuntimeValue)
			assert.Equal(suite.T(), test.rt, rtResult.TypeName)

			switch test.expected.(type) {
			case int:
				assert.Equal(suite.T(), test.expected, rtResult.Value)
			case float64:
				assert.Equal(suite.T(), test.expected, rtResult.Value)
			}

		})
	}

}

func (suite *TestInterpreterExpressionsTestSuite) Test_Comparison() {

	tests := []struct {
		input    string
		expected any
	}{
		{input: `
if testOption(1) == "hi" {
	fmt::println("we is hi")
}`, expected: 1},
		{input: `
var i int = 0
if testOption(i) == "hi" && testOption(i) == none {
	fmt::println("we is hi")
}`, expected: 1},
	}

	for _, test := range tests {
		suite.Run(test.input, func() {

			engine := NewTestingInterpreterEngine()
			engine.DisableTypeChecker = true
			engine.LoadSourceFromString(`
func testOption(i int) ?string {
  if i == 0 {
	return "hi"
  }
  return none
}

func eval() {
	` + test.input + `
}
`)
			engine.ProcessScripts()

			script := engine.SourceFiles[0]
			evalFunc := script.GetFunc("eval")

			w := utilities.NewIndentWriter(os.Stdout, " ")
			evalFunc.PrintTree(w.(*utilities.IndentWriter))

			assert.NotNil(suite.T(), evalFunc)

			call := &ast.CallExpression{
				AstNode:  ast.NewAstNode(nil),
				Function: ast.NewIdentifierWithValue(nil, "eval"),
			}

			rValue := engine.Evaluator.Eval(call)
			result := rValue.First()

			assert.IsType(suite.T(), new(ast.RuntimeValue), result)

			assert.NotNil(suite.T(), result)
			// rtResult := result.(*ast.RuntimeValue)

		})
	}

}
