package interpreter

import (
	"fmt"
	"os"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"testing"

	"arc/ast"
	"arc/lexer"
	"arc/utilities"
)

type TestInterpreterStatementsTestSuite struct {
	suite.Suite
}

func TestInterpreterStatementsSuite(t *testing.T) {
	suite.Run(t, new(TestInterpreterStatementsTestSuite))
}

func (suite *TestInterpreterStatementsTestSuite) SetupTest() {
	// Set an environment variable to indicate that we are running tests

	os.Setenv("INTERPRETER_TEST", "true")

}

func (suite *TestInterpreterStatementsTestSuite) TearDownSuite() {

}

func (suite *TestInterpreterStatementsTestSuite) Test_IfStatement() {

	scriptSrc := `

	fmt::printf("entry: %d\n", i);
    if i == 1 {
		fmt::printf("i is 1\n");
    } else if i == 2 {
		fmt::printf("i is 2\n");
    } else if i == 3 {
		fmt::printf("i is 3\n");
    } else {
		fmt::printf("i is not 1, 2, or 3\n");
    }
`
	tests := []struct {
		input string
		value int
	}{
		{input: scriptSrc, value: 1},
		{input: scriptSrc, value: 2},
		{input: scriptSrc, value: 3},
		{input: scriptSrc, value: 4},
	}

	for _, test := range tests {
		engine := NewTestingInterpreterEngine()
		scriptSrcFinal := `
func main() {
	var i = %d;
`
		scriptSrcFinal = fmt.Sprintf(scriptSrcFinal, test.value)
		scriptSrcFinal += scriptSrc
		scriptSrcFinal += "\n}"
		engine.LoadScriptFromString(scriptSrcFinal)
		engine.ProcessScripts()

		script := engine.SourceFiles[0]
		main := script.GetMainFunc()

		ifStatement := main.Body.Statements[2].(*ast.IfStatement)
		assert.Contains(suite.T(), ifStatement.GetTokenTypes(), lexer.TokenKeywordIf)

		assert.IsType(suite.T(), &ast.IfStatement{}, ifStatement.Else)
		assert.IsType(suite.T(), &ast.IfStatement{}, ifStatement.Else.(*ast.IfStatement).Else)
		assert.IsType(suite.T(), &ast.Block{}, ifStatement.Else.(*ast.IfStatement).Else.(*ast.IfStatement).Else)

		engine.Run()
	}
}

func (suite *TestInterpreterStatementsTestSuite) Test_WhileLoopStatement() {

	scriptSrc := `
func main() {
	var i = 0;
	for {
		fmt::printf("Looping: %d\n", i)
		if i == 10 {	
			break;
		}
		i++;
	}
	return i;
}
`

	engine := NewTestingInterpreterEngine()
	engine.LoadScriptFromString(scriptSrc)
	engine.ProcessScripts()

	script := engine.SourceFiles[0]
	main := script.GetMainFunc()

	whileLoop := main.Body.Statements[1].(*ast.LoopStatement)
	assert.Contains(suite.T(), whileLoop.GetTokenTypes(), lexer.TokenKeywordFor)
	assert.Nil(suite.T(), whileLoop.Range)
	assert.Nil(suite.T(), whileLoop.As)
	assert.Nil(suite.T(), whileLoop.Step)

	engine.Run()
}

func (suite *TestInterpreterStatementsTestSuite) Test_RangeLoopStatement() {

	scriptSrc := `
func main() {
	var i = 0;
	for 0..10 {
		fmt::printf("[loop - i] %d\n", i);
		fmt::printf("[loop - it] %d\n", it);
		i++;
	}
	return i;
}
`

	engine := NewTestingInterpreterEngine()
	engine.LoadScriptFromString(scriptSrc)
	engine.ProcessScripts()

	script := engine.SourceFiles[0]
	main := script.GetMainFunc()

	forLoop := main.Body.Statements[1].(*ast.LoopStatement)
	assert.Contains(suite.T(), forLoop.GetTokenTypes(), lexer.TokenKeywordFor)
	assert.Nil(suite.T(), forLoop.As)
	assert.Nil(suite.T(), forLoop.Step)
	assert.IsType(suite.T(), &ast.RangeExpression{}, forLoop.Range)
	// assert.Equal(suite.T(), 0, engine.Evaluator.MustEval(forLoop.Range.Left.(*ast.Literal)).(*ast.Literal).Value)
	// assert.Equal(suite.T(), 10, engine.Evaluator.MustEval(forLoop.Range.Right.(*ast.Literal)).(*ast.Literal).Value)

	engine.Run()

}

func (suite *TestInterpreterStatementsTestSuite) Test_RangeLoopStatementWithStep() {

	scriptSrc := `
func main() {
	var i = 0;
	for 1..10 step 2 {
		fmt::printf("[loop - i] %d\n", i);
		fmt::printf("[loop - it] %d\n", it);
		i += 2;
	}
	return i;
}
`

	engine := NewTestingInterpreterEngine()
	engine.LoadScriptFromString(scriptSrc)
	engine.ProcessScripts()

	script := engine.SourceFiles[0]
	main := script.GetMainFunc()

	loop := main.Body.Statements[1].(*ast.LoopStatement)
	assert.Contains(suite.T(), loop.GetTokenTypes(), lexer.TokenKeywordFor)
	assert.Nil(suite.T(), loop.As)
	assert.IsType(suite.T(), &ast.Literal{}, loop.Step)
	assert.IsType(suite.T(), &ast.RangeExpression{}, loop.Range)
	// assert.Equal(suite.T(), 1, engine.Evaluator.MustEval(loop.Range.Left.(*ast.Literal)).(*ast.Literal).Value)
	// assert.Equal(suite.T(), 10, engine.Evaluator.MustEval(loop.Range.Right.(*ast.Literal)).(*ast.Literal).Value)
	assert.Equal(suite.T(), 2, engine.Evaluator.MustEval(loop.Step.(*ast.Literal)).(*ast.Literal).Value)

	engine.Run()

}

func (suite *TestInterpreterStatementsTestSuite) Test_RangeLoopStatementWithStepWithAs() {

	scriptSrc := `
func main() {
	for 0..10 step 2 as i {
		fmt::printf("[loop - i] %d\n", i);
	}
}
`

	engine := NewTestingInterpreterEngine()
	engine.LoadScriptFromString(scriptSrc)
	engine.ProcessScripts()

	script := engine.SourceFiles[0]
	main := script.GetMainFunc()

	loop := main.Body.Statements[0].(*ast.LoopStatement)
	assert.Contains(suite.T(), loop.GetTokenTypes(), lexer.TokenKeywordFor)
	assert.IsType(suite.T(), &ast.Identifier{}, loop.As)
	assert.IsType(suite.T(), &ast.Literal{}, loop.Step)
	assert.IsType(suite.T(), &ast.RangeExpression{}, loop.Range)
	// assert.Equal(suite.T(), 0, engine.Evaluator.MustEval(loop.Range.Left.(*ast.Literal)).(*ast.Literal).Value)
	// assert.Equal(suite.T(), 10, engine.Evaluator.MustEval(loop.Range.Right.(*ast.Literal)).(*ast.Literal).Value)
	assert.Equal(suite.T(), 2, engine.Evaluator.MustEval(loop.Step.(*ast.Literal)).(*ast.Literal).Value)
	assert.Equal(suite.T(), "i", loop.As.Name)

	// Ensure that the step variable is bound to the environment correctly
	// assert.Nil(suite.T(), engine.Env.LookupVar("i"))
	// assert.Nil(suite.T(), engine.Env.LookupVar("rangeLower"))
	// assert.Nil(suite.T(), engine.Env.LookupVar("rangeUpper"))
	// engine.Evaluator.Eval(loop)
	//
	// // It should not be bound to our outer scope
	// assert.Nil(suite.T(), engine.Env.LookupVar("i"))
	// assert.Nil(suite.T(), engine.Env.LookupVar("rangeLower"))
	// assert.Nil(suite.T(), engine.Env.LookupVar("rangeUpper"))

	// It should be bound to the loop scope
	// assert.NotNil(suite.T(), loop.GetExecutionEnv().LookupVar("i"))
	// assert.NotNil(suite.T(), loop.GetExecutionEnv().LookupVar("rangeLower"))
	// assert.NotNil(suite.T(), loop.GetExecutionEnv().LookupVar("rangeUpper"))
	// assert.Equal(suite.T(), 0, loop.GetExecutionEnv().LookupVar("i"))
	// assert.Equal(suite.T(), 0, loop.GetExecutionEnv().LookupVar("rangeLower"))
	// assert.Equal(suite.T(), 10, loop.GetExecutionEnv().LookupVar("rangeUpper"))

	engine.Run()

}

func (suite *TestInterpreterStatementsTestSuite) Test_DeleteStatement() {

	scriptSrc := `
func main() {

	fmt::printf("Dictionary delete\n");
	var d = {
		"a": "foo",
		"b": "bar",
		"c": "baz",
	};
	delete d["a"];
	delete d.b;
	var c = "c";
	delete d[c];
	fmt::printf("Dictionary end result: %v\n", d);

	fmt::printf("Array delete\n");
	var a []int = {1, 2, 3, 4, 5, 6, 7};
	delete a[0];
	delete a[1];
	var i = 0;
	delete a[i];
	delete a[0:];

	fmt::printf("Array end result: %v\n", a);

}
`

	engine := NewTestingInterpreterEngine()
	engine.LoadScriptFromString(scriptSrc)
	engine.ProcessScripts()

	script := engine.SourceFiles[0]
	main := script.GetMainFunc()

	timer := utilities.NewTimer("Eval")
	engine.Evaluator.Eval(main.Body)
	timer.StopAndLog()

}
