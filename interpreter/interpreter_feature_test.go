package interpreter

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"testing"

	"arc/utilities"
)

type TestInterpreterFeatureTestSuite struct {
	suite.Suite
}

func TestInterpreterFeatureSuite(t *testing.T) {
	suite.Run(t, new(TestInterpreterFeatureTestSuite))
}

func (suite *TestInterpreterFeatureTestSuite) SetupTest() {

}

func (suite *TestInterpreterFeatureTestSuite) TearDownSuite() {

}

func (suite *TestInterpreterFeatureTestSuite) Test_Arrays() {
	scriptSrc := `
object User {
   name string
}
func main() {
	var user []User = {
	    User{name:"John"},
	    User{name:"Bruce"},
	};
	var johns = user[0:1];
	var bruces = user[1:2];
	var all = user[0:];

	fmt::printf("Johns: %v\n", johns);
	fmt::printf("Bruces: %v\n", bruces);
	fmt::printf("All: %v\n", all);
	fmt::printf("John: %v\n", user[0]);
	fmt::printf("Bruce: %v\n", user[1]);
}
`

	engine := NewTestingInterpreterEngine()
	engine.LoadSourceFromString(scriptSrc)
	engine.ProcessScripts()

	script := engine.SourceFiles[0]
	main := script.GetMainFunc()

	engine.Evaluator.Eval(main.Body)

	userVar := engine.Env.LookupVar("user")
	assert.NotNil(suite.T(), userVar, "user variable not found")
}

func (suite *TestInterpreterFeatureTestSuite) Test_ArrayLoop() {
	scriptSrc := `
object User {
   name string
}
func main() {
	var users []User = {
	    User{name:"John"},
	    User{name:"Bruce"},
	};

  	for users {
		fmt::printf("User: %v\n", it);
	}

  	for users as u {
		fmt::printf("User(u): %v\n", u);
	}

}
`

	engine := NewTestingInterpreterEngine()
	engine.LoadSourceFromString(scriptSrc)
	engine.ProcessScripts()

	script := engine.SourceFiles[0]
	main := script.GetMainFunc()

	engine.Evaluator.Eval(main.Body)

	userVar := engine.Env.LookupVar("user")
	assert.NotNil(suite.T(), userVar, "user variable not found")
}

func (suite *TestInterpreterFeatureTestSuite) Test_Dictionary() {
	scriptSrc := `
func main() {
	var value = {
		"key1": "value1",
		"key2": "value2",
	};

	var key1 = "key1";
	var key2 = "key2";

	fmt::printf("Value: %v\n", value);
	fmt::printf("Key1: %v\n", value.key1);
	fmt::printf("Key2: %v\n", value.key2);
	fmt::printf("Key1: %v\n", value["key1"]);
	fmt::printf("Key2: %v\n", value["key2"]);

	var key1 = "key1";
	var key2 = "key2";
	fmt::printf("Key1: %v\n", value[key1]);
	fmt::printf("Key2: %v\n", value[key2]);

	value.key1 = "value1.1";
	value.key2 = "value2.1";
	
	value["key3"] = "value3";
	fmt::printf("Value: %v\n", value);

	value["key3"] = "value3.1";
	fmt::printf("Value: %v\n", value);

	var key4 = "key4";
	value[key4] = "value4";

	value.key5 = "value5";

	fmt::printf("Value: %v\n", value);
	
	delete value["key3"];
	delete value[key4];
	delete value.key5;

	fmt::printf("Value: %v\n", value);

	delete value;

	fmt::printf("Value: %v\n", value);
}
`

	engine := NewTestingInterpreterEngine()
	engine.LoadSourceFromString(scriptSrc)
	engine.ProcessScripts()

	script := engine.SourceFiles[0]
	main := script.GetMainFunc()

	timer := utilities.NewTimer("Eval")
	engine.Evaluator.Eval(main.Body)
	timer.StopAndLog()

	value := engine.Env.LookupVar("value")
	assert.NotNil(suite.T(), value, "value variable not found")
}
