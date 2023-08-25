

## Routes

**All routes must be defined in a `http` block**

Example:
```arc

http {
    route GET "/hello" {
        return text "Hello world"
    }
    // GET|PUT|POST|PATCH|DELETE|HEAD|OPTIONS are valid methods
}
```

## Body

### "Injecting" into a variable

This will convert the json body into a variable of the given type.

```arc
object User {
   name string
}
http {
    route POST "/data" {
        from body as user User;
    }
}
```

## Responses:

```arc 
return <json|text|html> <value> <status <number>?>
```

Examples:
```arc
return { "error": "Something went wrong!" } status 500
return status 404
return text "Hello world"

object User {
   name string
}
var user User = User{ name: "John" }
return user;
```
