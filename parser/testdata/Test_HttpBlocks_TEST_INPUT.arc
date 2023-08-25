http {
    route POST "/data/json" {
        from body as user User;

        return text "Hello world"
        return user;

        return { "error": "Something went wrong!" } status 500

        return text "Not found" status 404
        return status 404

    }
}

