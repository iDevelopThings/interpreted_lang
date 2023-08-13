object User {
   body string
   request string
}

route POST "/data" {
    from body as user User;

    var message = request.body.message;
	fmt::printf("\n\n");
	fmt::printf("RequestBody: %v\n", request.body);
	fmt::printf("Message: %s\n", message);
	fmt::printf("\n\n");
}
