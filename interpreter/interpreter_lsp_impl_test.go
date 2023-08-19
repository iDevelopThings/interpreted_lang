package interpreter

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	protocol "github.com/tliron/glsp/protocol_3_16"

	"testing"
)

type TestInterpreterLspImplTestSuite struct {
	suite.Suite
}

func TestInterpreterLspImplSuite(t *testing.T) {
	suite.Run(t, new(TestInterpreterLspImplTestSuite))
}

func (suite *TestInterpreterLspImplTestSuite) SetupTest() {

}

func (suite *TestInterpreterLspImplTestSuite) TearDownSuite() {

}

func (suite *TestInterpreterLspImplTestSuite) Test_TypeChecker_GetNodeAtPosition() {
	engine := NewTestingInterpreterEngine()
	engine.LoadScript("test_data/testing.sl")
	engine.ProcessScripts()

	node := TypeChecker.GetNodeAtPosition(
		engine.SourceFiles[0].Program,
		protocol.Position{
			Line:      15,
			Character: 17,
		},
	)
	assert.NotNil(suite.T(), node)

	nodeType, nodeTypeDeclaration := TypeChecker.FindDeclaration(node)

	assert.NotNil(suite.T(), nodeType)
	assert.NotNil(suite.T(), nodeTypeDeclaration)

	nodeTypeDefinitionSource := TypeChecker.GetNodeSourceFile(nodeType)
	assert.NotNil(suite.T(), nodeTypeDefinitionSource)

	nodeTypePosRange := nodeTypeDeclaration.GetTokenRange()
	nodeTypePosRange.ZeroIndexed()

	nodeDeclTypePosRange := nodeType.GetTokenRange()
	nodeDeclTypePosRange.ZeroIndexed()

	assert.NotNil(suite.T(), nodeTypePosRange)
	assert.NotNil(suite.T(), nodeDeclTypePosRange)

	response := []*protocol.Location{
		{
			URI: nodeTypeDefinitionSource.Path,
			Range: protocol.Range{
				Start: protocol.Position{
					Line:      protocol.UInteger(nodeDeclTypePosRange.StartLine),
					Character: protocol.UInteger(nodeDeclTypePosRange.StartCol),
				},
				End: protocol.Position{
					Line:      protocol.UInteger(nodeDeclTypePosRange.StopLine),
					Character: protocol.UInteger(nodeDeclTypePosRange.RangeStopCol),
				},
			},
		},
		{
			URI: nodeTypeDefinitionSource.Path,
			Range: protocol.Range{
				Start: protocol.Position{
					Line:      protocol.UInteger(nodeTypePosRange.StartLine),
					Character: protocol.UInteger(nodeTypePosRange.StartCol),
				},
				End: protocol.Position{
					Line:      protocol.UInteger(nodeTypePosRange.StopLine),
					Character: protocol.UInteger(nodeTypePosRange.RangeStopCol),
				},
			},
		},
	}

	assert.NotNil(suite.T(), response)
}
