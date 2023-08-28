package lexer

import (
	"flag"
	"os"
	"testing"

	"github.com/charmbracelet/log"
	"github.com/goccy/go-json"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/suite"
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

//	func loadGolden[T any](t *testing.T, name string, data T) T {
//		t.Helper()
//
//		goldenFileName := "testdata/" + name + "_TEST.golden"
//		content, err := os.ReadFile(goldenFileName)
//		if err != nil {
//			if os.IsNotExist(err) {
//				if *update {
//					writeGolden(t, name, data)
//					return data
//				}
//				t.Fatalf("golden file %q does not exist, run with -update to create it", goldenFileName)
//			}
//
//			t.Fatalf("failed to load golden file %q: %v", goldenFileName, err)
//		}
//
//		var v T
//		err = json.Unmarshal(content, &v)
//		if err != nil {
//			t.Fatalf("failed to unmarshal golden file %q: %v", goldenFileName, err)
//		}
//
//		if *update {
//			writeGolden(t, name, data)
//		}
//
//		return v
//	}
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

	err = os.WriteFile(goldenFileName, jsonData, 0600)
	if err != nil {
		t.Fatalf("failed to write golden file %q: %v", goldenFileName, err)
	}
}

type TestLexerTestSuite struct {
	suite.Suite
	rawGoldenData string
	goldenData    []*Token
	inputData     string
}

func TestLexerSuite(t *testing.T) {
	suite.Run(t, new(TestLexerTestSuite))
}

func (suite *TestLexerTestSuite) SetupTest() {
	log.SetLevel(log.DebugLevel)
}

func (suite *TestLexerTestSuite) TearDownSuite() {

}
func (suite *TestLexerTestSuite) BeforeTest(suiteName, testName string) {
	goldenData, loaded := loadGoldenWithoutTestData(suite.T(), testName)
	if loaded {
		suite.rawGoldenData = goldenData
		if err := json.Unmarshal([]byte(goldenData), &suite.goldenData); err != nil {
			suite.Errorf(err, "failed to unmarshal golden data")
		}
	}

	inputFileName := testName + "_TEST_INPUT.arc"

	suite.inputData, _ = loadFileContent(suite.T(), inputFileName)
}
func (suite *TestLexerTestSuite) AfterTest(suiteName, testName string) {
	if *update {
		writeGolden(suite.T(), testName, suite.goldenData)
		suite.T().Logf("updated golden file for %s", testName)
	}
}

func (suite *TestLexerTestSuite) HandleStats(suiteName string, stats *suite.SuiteInformation) {
	suite.T().Logf("TestSuite %s finished in %s", suiteName, stats.End.Sub(stats.Start))
}

func checkForUnknownTokens(t *testing.T, tokens []*Token) {
	t.Helper()

	for _, tok := range tokens {
		if tok.Is(TokenUnknown) {
			t.Errorf("unknown token %q at %s", tok.Value, tok.Pos)
		}
	}
}

func (suite *TestLexerTestSuite) Test_LexerInput() {
	l := NewLexer(suite.inputData)
	tokens := l.readAll()

	checkForUnknownTokens(suite.T(), tokens)

	if diff := cmp.Diff(suite.goldenData, tokens); diff != "" {
		suite.Failf("Token mismatch (-want +got):\n%s", diff)
	}

	suite.goldenData = tokens

	// l.debugDisplayTokens(suite.T().Logf, tokens)
}

func (suite *TestLexerTestSuite) Test_LexerInput_MassInput() {
	l := NewLexer(suite.inputData)
	tokens := l.readAll()
	checkForUnknownTokens(suite.T(), tokens)

	if diff := cmp.Diff(suite.goldenData, tokens); diff != "" {
		suite.Failf("Token mismatch (-want +got):\n%s", diff)
	}

	suite.goldenData = tokens

	// l.debugDisplayTokens(suite.T().Logf, tokens)
}

func (suite *TestLexerTestSuite) Test_LexerInput_Comments() {
	l := NewLexer(suite.inputData)
	tokens := l.readAll()
	checkForUnknownTokens(suite.T(), tokens)

	if diff := cmp.Diff(suite.goldenData, tokens); diff != "" {
		suite.Failf("Token mismatch (-want +got):\n%s", diff)
	}

	suite.goldenData = tokens

	// l.debugDisplayTokens(suite.T().Logf, tokens)
}

func (suite *TestLexerTestSuite) Test_Identifiers() {
	inputStr := ""
	var expectedTokens []*Token
	pos := &Position{Line: 1, Column: 0, Abs: 0}

	for _, match := range keywordMatchTable {
		inputStr += match.Value + " "
		pos.Column += len(match.Value)
		pos.Abs += len(match.Value)

		tok := NewToken(match.Value, TokenIdentifier, match.Token)
		tok.Pos = NewTokenPosition(&Position{
			pos.Line,
			pos.Column,
			pos.Abs,
		}, len(match.Value))

		pos.Column += 1
		pos.Abs += 1
		expectedTokens = append(expectedTokens, tok)
	}

	l := NewLexer(inputStr)
	tokens := l.readAll()
	checkForUnknownTokens(suite.T(), tokens)

	if diff := cmp.Diff(expectedTokens, tokens); diff != "" {
		suite.T().Logf("Input: %q", inputStr)
		suite.Failf("Token mismatch (-want +got):\n%s", diff)
	}
}

func (suite *TestLexerTestSuite) Test_Range() {
	l := NewLexer(suite.inputData)
	tokens := l.readAll()
	checkForUnknownTokens(suite.T(), tokens)

	if diff := cmp.Diff(suite.goldenData, tokens); diff != "" {
		suite.Failf("Token mismatch (-want +got):\n%s", diff)
	}

	suite.goldenData = tokens

	l.debugDisplayTokens(suite.T().Logf, tokens)
}

func (suite *TestLexerTestSuite) Test_ArgList() {
	l := NewLexer(suite.inputData)
	tokens := l.readAll()
	checkForUnknownTokens(suite.T(), tokens)

	if diff := cmp.Diff(suite.goldenData, tokens); diff != "" {
		suite.Failf("Token mismatch (-want +got):\n%s", diff)
	}

	suite.goldenData = tokens

	l.debugDisplayTokens(suite.T().Logf, tokens)
}
func (suite *TestLexerTestSuite) Test_Casing() {
	l := NewLexer(suite.inputData)
	tokens := l.readAll()
	checkForUnknownTokens(suite.T(), tokens)

	if diff := cmp.Diff(suite.goldenData, tokens); diff != "" {
		suite.Failf("Token mismatch (-want +got):\n%s", diff)
	}

	suite.goldenData = tokens

	l.debugDisplayTokens(suite.T().Logf, tokens)
}
