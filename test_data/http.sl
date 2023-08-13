object Account {
	isAdmin bool
}
object User {
   name string
   account Account
}

route POST "/data/json" {
	from body as user User;

	fmt::printf("\n\n");
	fmt::printf("RequestBody: %v\n", user);
	fmt::printf("Name: %s\n", user.name);
	fmt::printf("IsAdmin: %v\n", user.account.isAdmin);
	fmt::printf("\n\n");

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

func main(){}

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
