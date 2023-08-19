import "imports.sl"
import "dictionaries.sl"
import "http_basic_test.sl"

func pls() {
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
}

func main() {
    pls()
}

object Account {
	isAdmin bool
}
object User {
   name string
   account Account
}

route POST "/data/json" {
	from body as user User;

    var testing = "..";

	// fmt::printf("\n\n");
	// fmt::printf("RequestBody: %v\n", user);
	// fmt::printf("Name: %s\n", user.name);
	// fmt::printf("IsAdmin: %v\n", user.account.isAdmin);
	// fmt::printf("\n\n");

	respond with user;
}

route POST "/data/form" {
	//from body as user User;
    //fmt::printf("\n\n");
    //fmt::printf("RequestBody: %v\n", user);
    //fmt::printf("Name: %s\n", user.name);
    //fmt::printf("IsAdmin: %v\n", user.account.isAdmin);
    //fmt::printf("\n\n");
    //respond with user;
}
route POST "/data/multipart_form" {
	//from body as user User;
	//fmt::printf("\n\n");
	//fmt::printf("RequestBody: %v\n", user);
	//fmt::printf("Name: %s\n", user.name);
	//fmt::printf("IsAdmin: %v\n", user.account.isAdmin);
	//fmt::printf("\n\n");
	//respond with user;
}


// route GET "/" {
//     respond with text "Hello world"
// }
// route GET "/notfound" {
//     respond with text "Not found" status 404
// }
// route GET "/notfound_2" {
//     respond with status 404
// }
//
// route GET "/error" {
//     respond with { "error": "Something went wrong!" } status 500
// }
//
// route POST "/data" {
//     var name = request.body["name"];
// }

// middleware "/admin/*" {
//     // Some kind of authentication
//     if not authenticated(request) {
//         response 401 "Unauthorized"
//         return
//     }
// }


object MyObject {
    StrValue string
    IntValue int
    Nested MyObject*
}

func (obj MyObject) TestFunc() int {
    // fmtt::fmt::printf("testing func: %s\n", obj.StrValue)
    var a = 1 + 2 + obj.IntValue;

    return a;
}

func runInputScript() {
    var z MyObject* = MyObject{
        StrValue: "root str",
        IntValue: (1 + 2) * 3,
    };

    fmt::printf("testing func: %d\n", z.TestFunc());
    // a.b.thing();

    // var a = 1 + 2;
    // var b = 1 + 2 ^ -3;
    // var z MyObject* = MyObject{
    //     StrValue: "root str",
    //     Nested: MyObject{
    //         StrValue: "nested str",
    //         IntValue: 123,
    //     },
    // };
    // fmt::printf("testing var: %s - nested: %s\n", z.StrValue, z.Nested.StrValue)
}


object User {
   name string
}
func (u User) String() string {
    return u.name;
}

func pls() {
	var users []User = {
	    User{name:"John"},
	    User{name:"Bruce"},
	};


	fmt::printf("Len: %s\n", users.length());

    var barry = User{name:"Barry"};
    users.add(barry);

	fmt::printf("Len: %s\n", users.length());

    users.remove(0);
	users.remove(barry);

	fmt::printf("Len: %s\n", users.length());

  	for users {
		fmt::printf("User: %v\n", it);
	}

  	for users as u {
		fmt::printf("User(u): %v\n", u);
	}
}