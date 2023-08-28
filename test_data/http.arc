object SomeData {
    message string
    number int
}

http {
    route GET "/hello" {
        return text "Hello world"
    }

    route GET "/hello/dict" {
        return { "hello": "world" } status 201
    }
    route GET "/hello/404" {
        return text "Not found!!!!" status 404
    }
    route GET "/hello/500/dict" {
        return { "error": "Something went wrong!!!" } status 500
    }

    route GET "/injections/body" {
        from body as data SomeData
//        data.message += "pls sir"
        return {"msg": data}
    }

    route GET "/injections/param/:name" {
        from body as data SomeData
        from route as name string
//        data.message += " - name is " + name
        return {"data" : data}
    }

    route GET "/injections/query" {
        from body as data SomeData
        from query as name string
//        data.message += " - name is " + name
        return {"data" : data}
    }

    route POST "/data/json" {
        from body as data SomeData

        return text "Hello world"
        return user


        return text "Not found" status 404
        return status 404
    }
}

func main() {}














// import "imports.sl"
// import "dictionaries.sl"
// import "http_basic_test.sl"
// object Account {
// 	isAdmin bool
// }
// object User {
//    name string
//    account Account
// }
//
// http {
//     route POST "/data/json" {
//         from body as user User;
//
//         var testing = "..";
//
//
//         // fmt::printf("\n\n");
//         // fmt::printf("RequestBody: %v\n", user);
//         // fmt::printf("Name: %s\n", user.name);
//         // fmt::printf("IsAdmin: %v\n", user.account.isAdmin);
//         // fmt::printf("\n\n");
//
//         return text "Hello world"
//         return text "Not found" status 404
//         return status 404
//         return { "error": "Something went wrong!" } status 500
//         return user;
//
//         respond with text "Hello world"
//         respond with text "Not found" status 404
//         respond with status 404
//         respond with { "error": "Something went wrong!" } status 500
//         respond with user;
//     }
// }
//
//
//
// route POST "/data/form" {
// 	//from body as user User;
//     //fmt::printf("\n\n");
//     //fmt::printf("RequestBody: %v\n", user);
//     //fmt::printf("Name: %s\n", user.name);
//     //fmt::printf("IsAdmin: %v\n", user.account.isAdmin);
//     //fmt::printf("\n\n");
//     //respond with user;
// }
// route POST "/data/multipart_form" {
// 	//from body as user User;
// 	//fmt::printf("\n\n");
// 	//fmt::printf("RequestBody: %v\n", user);
// 	//fmt::printf("Name: %s\n", user.name);
// 	//fmt::printf("IsAdmin: %v\n", user.account.isAdmin);
// 	//fmt::printf("\n\n");
// 	//respond with user;
// }
//
//
// func main(){
//
// }

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
