package interpreter

import (
	"net/http/httptest"
	"net/url"
	"strings"

	"github.com/stretchr/testify/suite"

	"testing"

	"arc/http_server"
	"arc/utilities"
)

type TestInterpreterHttpTestSuite struct {
	suite.Suite
}

func TestInterpreterHttpSuite(t *testing.T) {
	suite.Run(t, new(TestInterpreterHttpTestSuite))
}

func (suite *TestInterpreterHttpTestSuite) SetupTest() {

}

func (suite *TestInterpreterHttpTestSuite) TearDownSuite() {

}

func (suite *TestInterpreterHttpTestSuite) Test_PostMethodWithDataInjection() {

	data := map[string]any{
		"name": "John",
		"account": map[string]any{
			"isAdmin": true,
		},
	}

	script := `
object Account {
	isAdmin bool
}
object User {
   name string
   account Account
}

route POST "/data" {
	from body as user User;

	fmt::printf("\n\n");
	fmt::printf("RequestBody: %v\n", user);
	fmt::printf("Name: %s\n", user.name);
	fmt::printf("IsAdmin: %v\n", user.account.isAdmin);
	fmt::printf("\n\n");
}`

	engine := NewTestingInterpreterEngine()
	engine.LoadScriptFromString(script)
	engine.Run()

	w := httptest.NewRecorder()
	r := httptest.NewRequest(
		"POST",
		"/data",
		strings.NewReader(string(utilities.ToJson(data))),
	)
	r.Header.Set("Content-Type", "application/json")

	timer := utilities.NewTimer("Processing Time")
	http_server.GetRouter().ServeHTTP(w, r)
	timer.StopAndLog()

	suite.Equal(200, w.Code)

}

func (suite *TestInterpreterHttpTestSuite) Test_RequestBodyAccess() {

	data := map[string]any{
		"name": "John",
		"account": map[string]any{
			"isAdmin": true,
		},
	}

	script := `
route POST "/data" {
	fmt::printf("name: %s\n", request.body.name);
	fmt::printf("name[]: %s\n", request.body["name"]);

	fmt::printf("isAdmin: %v\n", request.body.account.isAdmin);
	fmt::printf("isAdmin[]: %v\n", request.body.account["isAdmin"]);
}`

	engine := NewTestingInterpreterEngine()
	engine.LoadScriptFromString(script)
	engine.Run()

	w := httptest.NewRecorder()
	r := httptest.NewRequest(
		"POST",
		"/data",
		strings.NewReader(string(utilities.ToJson(data))),
	)
	r.Header.Set("Content-Type", "application/json")

	timer := utilities.NewTimer("Processing Time")
	http_server.GetRouter().ServeHTTP(w, r)
	timer.StopAndLog()

	suite.Equal(200, w.Code)

}

func (suite *TestInterpreterHttpTestSuite) Test_RequestFormUrlEncoded() {
	data := url.Values{}
	data.Add("name[]", "John")
	data.Add("name[]", "Doe")
	data.Add("account[isAdmin]", "true")

	script := `
object Account {
	   isAdmin bool
}
object User {
	   name []string
	   account Account
}
route POST "/data" {
	from body as user User;

	fmt::printf("name: %s\n", request.body.name);
	fmt::printf("name[]: %s\n", request.body["name"]);

	fmt::printf("isAdmin: %v\n", request.body.account.isAdmin);
	fmt::printf("isAdmin[]: %v\n", request.body.account["isAdmin"]);
	
	fmt::printf("user.name: %s\n", user.name);
	fmt::printf("user.isAdmin: %v\n", user.account.isAdmin);
}`

	engine := NewTestingInterpreterEngine()
	engine.LoadScriptFromString(script)
	engine.Run()

	w := httptest.NewRecorder()
	r := httptest.NewRequest(
		"POST",
		"/data",
		strings.NewReader(data.Encode()),
	)
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	timer := utilities.NewTimer("Processing Time")
	http_server.GetRouter().ServeHTTP(w, r)
	timer.StopAndLog()

	suite.Equal(200, w.Code)

}
