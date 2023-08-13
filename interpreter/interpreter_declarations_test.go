package interpreter

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"testing"

	"interpreted_lang/ast"
)

type TestInterpreterDeclarationTestSuite struct {
	suite.Suite
}

func TestInterpreterDeclarationSuite(t *testing.T) {
	suite.Run(t, new(TestInterpreterDeclarationTestSuite))
}

func (suite *TestInterpreterDeclarationTestSuite) SetupTest() {

}

func (suite *TestInterpreterDeclarationTestSuite) TearDownSuite() {

}

func (suite *TestInterpreterDeclarationTestSuite) Test_Objects() {
	scriptSrc := `
object Account {
	isAdmin bool
}
object User {
   name string
   account Account
}
func main() {
	var user = User{ name : "John", account : Account{ isAdmin : true } };
}
`

	engine := NewTestingInterpreterEngine()
	engine.LoadScriptFromString(scriptSrc)
	engine.prepareToEvaluate()

	script := engine.Scripts[0]
	main := script.GetMainFunc()

	engine.Evaluator.Eval(main.Body)

	userVar := engine.Env.LookupVar("user")
	assert.NotNil(suite.T(), userVar, "user variable not found")
	assert.IsType(suite.T(), &ast.RuntimeValue{}, userVar, "user variable is not an object")

	userObj := userVar.(*ast.RuntimeValue)
	assert.Equal(suite.T(), "User", userObj.TypeName, "user variable is not of type User")
	assert.Equal(suite.T(), "John", userObj.GetField("name"), "user.name is not John")
	assert.Equal(suite.T(), true, userObj.GetField("account").GetField("isAdmin"), "user.account.isAdmin is not true")

}

func (suite *TestInterpreterDeclarationTestSuite) Test_ObjectMemberAccess() {
	scriptSrc := `
object Account {
	isAdmin bool
}
object User {
   name string
   account Account
}
func main() {
	var user = User{ name : "John", account : Account{ isAdmin : true } };

	var name = user.name;
	var isAdmin = user.account.isAdmin;
}
`

	engine := NewTestingInterpreterEngine()
	engine.LoadScriptFromString(scriptSrc)
	engine.prepareToEvaluate()

	script := engine.Scripts[0]
	main := script.GetMainFunc()

	engine.Evaluator.Eval(main.Body)

	userVar := engine.Env.LookupVar("user")
	assert.NotNil(suite.T(), userVar, "user variable not found")
	assert.IsType(suite.T(), &ast.RuntimeValue{}, userVar, "user variable is not a runtime value")

	userObj := userVar.(*ast.RuntimeValue)
	assert.Equal(suite.T(), "User", userObj.TypeName, "user variable is not of type User")
	assert.Equal(suite.T(), "John", userObj.GetField("name"), "user.name is not John")
	assert.Equal(suite.T(), true, userObj.GetField("account").GetField("isAdmin"), "user.account.isAdmin is not true")

	nameVar := engine.Env.LookupVar("name")
	assert.NotNil(suite.T(), nameVar, "name variable not found")

	isAdminVar := engine.Env.LookupVar("isAdmin")
	assert.NotNil(suite.T(), isAdminVar, "isAdmin variable not found")
}

func (suite *TestInterpreterDeclarationTestSuite) Test_ObjectMethods() {
	scriptSrc := `
object User {
   name string
   account Account
}

func (u User) GetName() string {
 	return u.name;
}
func (u User) SetName(name string) {
 	u.name = name;
}

func main() {
	var user = User{ name : "John" };

	var name = user.GetName();
    fmt::printf("Name: %s\n", name);
	user.SetName("Bruce");
	name = user.GetName();
	fmt::printf("Name: %s\n", name);
}
`

	engine := NewTestingInterpreterEngine()
	engine.LoadScriptFromString(scriptSrc)
	engine.prepareToEvaluate()

	script := engine.Scripts[0]
	main := script.GetMainFunc()

	engine.Evaluator.Eval(main.Body)

	userVar := engine.Env.LookupVar("user")
	assert.NotNil(suite.T(), userVar, "user variable not found")
	assert.IsType(suite.T(), &ast.RuntimeValue{}, userVar, "user variable is not an object")

	userObj := userVar.(*ast.RuntimeValue)
	assert.Equal(suite.T(), "User", userObj.TypeName, "user variable is not of type User")
	assert.Equal(suite.T(), "Bruce", userObj.GetField("name").Value, "user.name is not John")

	nameVar := engine.Env.LookupVar("name")
	assert.NotNil(suite.T(), nameVar, "name variable not found")

}
