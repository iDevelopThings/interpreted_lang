package parser

import (
	"encoding/json"
	"flag"
	"os"
	"testing"

	"github.com/charmbracelet/log"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/suite"

	"arc/ast"
	"arc/lexer"
	"arc/utilities"
)

var update = flag.Bool("update", false, "update the .golden files")

func loadFileContent(t *testing.T, name string) (string, bool) {
	t.Helper()

	fileName := "testdata/" + name
	content, err := os.ReadFile(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return "", false
		}
		t.Fatalf("failed to load file %q: %v", fileName, err)
	}

	return string(content), true
}
func loadGolden[T any](t *testing.T, name string, data T) T {
	t.Helper()

	goldenFileName := "testdata/" + name + "_TEST.golden"
	content, err := os.ReadFile(goldenFileName)
	if err != nil {
		if os.IsNotExist(err) {
			if *update {
				writeGolden(t, name, data)
				return data
			}
			t.Fatalf("golden file %q does not exist, run with -update to create it", goldenFileName)
		}

		t.Fatalf("failed to load golden file %q: %v", goldenFileName, err)
	}

	var v T
	err = json.Unmarshal(content, &v)
	if err != nil {
		t.Fatalf("failed to unmarshal golden file %q: %v", goldenFileName, err)
	}

	if *update {
		writeGolden(t, name, data)
	}

	return v
}
func loadGoldenWithoutTestData(t *testing.T, name string) (string, bool) {
	t.Helper()

	goldenFileName := "testdata/" + name + "_TEST.golden"
	content, err := os.ReadFile(goldenFileName)
	if err != nil {
		if os.IsNotExist(err) {
			return "", false
		}
		t.Fatalf("failed to load golden file %q: %v", goldenFileName, err)
	}

	return string(content), true
}
func writeGolden(t *testing.T, name string, testData any) {
	t.Helper()

	goldenFileName := "testdata/" + name + "_TEST.golden"

	jsonData, err := json.MarshalIndent(testData, "", "  ")
	if err != nil {
		t.Fatalf("failed to marshal test data: %v", err)
	}

	err = os.WriteFile(goldenFileName, jsonData, 0644)
	if err != nil {
		t.Fatalf("failed to write golden file %q: %v", goldenFileName, err)
	}
}

type TestParserTestSuite struct {
	suite.Suite
	rawGoldenData string
	goldenData    *ast.Program
	inputData     string
}

func TestLexerSuite(t *testing.T) {
	suite.Run(t, new(TestParserTestSuite))
}

func (suite *TestParserTestSuite) SetupTest() {
	log.SetLevel(log.DebugLevel)
}

func (suite *TestParserTestSuite) TearDownSuite() {

}
func (suite *TestParserTestSuite) BeforeTest(suiteName, testName string) {
	// goldenData, loaded := loadGoldenWithoutTestData(suite.T(), testName)
	// if loaded {
	// 	suite.rawGoldenData = goldenData
	// 	if err := json.Unmarshal([]byte(goldenData), &suite.goldenData); err != nil {
	// 		suite.Failf("failed to unmarshal golden data", err.Error())
	// 	}
	// }
	//
	inputFileName := testName + "_TEST_INPUT.arc"
	suite.inputData, _ = loadFileContent(suite.T(), inputFileName)
}
func (suite *TestParserTestSuite) AfterTest(suiteName, testName string) {
	// if *update {
	// 	writeGolden(suite.T(), testName, suite.goldenData)
	// 	suite.T().Logf("updated golden file for %s", testName)
	// }
}

func (suite *TestParserTestSuite) Test_ParserInput() {
	l := lexer.NewLexer(suite.inputData)
	p := NewParser(l)

	program := p.Parse()
	w := utilities.NewIndentWriter(os.Stdout, " ")

	program.PrintTree(w.(*utilities.IndentWriter))

	// if diff := cmp.Diff(suite.goldenData, tokens); diff != "" {
	// 	suite.Failf("Token mismatch (-want +got):\n%s", diff)
	// }
	// suite.goldenData = tokens
}

func (suite *TestParserTestSuite) Test_TokenGrouping() {

	l := lexer.NewLexer(suite.inputData)
	p := NewParser(l)

	program := p.Parse()

	// program.Functions[0].GetTokenBounds()

	if diff := cmp.Diff(suite.goldenData, program); diff != "" {
		suite.Failf("Token mismatch (-want +got):\n%s", diff)
	}
	suite.goldenData = program

	w := utilities.NewIndentWriter(os.Stdout, " ")
	program.PrintTree(w.(*utilities.IndentWriter))
}

func (suite *TestParserTestSuite) Test_Enums() {

	l := lexer.NewLexer(suite.inputData)
	p := NewParser(l)
	program := p.Parse()

	if diff := cmp.Diff(suite.goldenData, program); diff != "" {
		suite.Failf("Token mismatch (-want +got):\n%s", diff)
	}
	suite.goldenData = program

	w := utilities.NewIndentWriter(os.Stdout, " ")
	program.PrintTree(w.(*utilities.IndentWriter))
}

func (suite *TestParserTestSuite) Test_Defer() {

	l := lexer.NewLexer(suite.inputData)
	p := NewParser(l)
	program := p.Parse()

	if diff := cmp.Diff(suite.goldenData, program); diff != "" {
		suite.Failf("Token mismatch (-want +got):\n%s", diff)
	}
	suite.goldenData = program

	w := utilities.NewIndentWriter(os.Stdout, " ")
	program.PrintTree(w.(*utilities.IndentWriter))
}
func (suite *TestParserTestSuite) Test_Errors() {
	l := lexer.NewLexer(suite.inputData)
	l.SetSource("testdata/Test_Errors_TEST_INPUT.arc")
	p := NewParser(l)
	program := p.Parse()

	if diff := cmp.Diff(suite.goldenData, program); diff != "" {
		suite.Failf("Token mismatch (-want +got):\n%s", diff)
	}
	suite.goldenData = program

	w := utilities.NewIndentWriter(os.Stdout, " ")
	program.PrintTree(w.(*utilities.IndentWriter))
}

func (suite *TestParserTestSuite) Test_GeneratedVisitor() {
	data, _ := loadFileContent(suite.T(), "Test_Errors_TEST_INPUT.arc")
	l := lexer.NewLexer(data)
	l.SetSource("testdata/Test_Errors_TEST_INPUT.arc")
	p := NewParser(l)
	program := p.Parse()

	ast.Walk(program, func(node ast.Node) any {
		switch node := node.(type) {
		case *ast.FunctionDeclaration:
			suite.T().Logf("FunctionDeclaration: %#v", node)
		case *ast.CallExpression:
			suite.T().Logf("CallExpression: %#v", node)
		}

		// suite.T().Logf("node: %#v", node)

		return true
	})

	// w := utilities.NewIndentWriter(os.Stdout, " ")
	// program.PrintTree(w.(*utilities.IndentWriter))
}
