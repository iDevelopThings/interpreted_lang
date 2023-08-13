

## Routes

```sl
route <method> "<path>" {
    <handler>
}
```

Example:

```sl
route POST "/data" {
}
```

## Body

### "Injecting" into a variable

This will convert the json body into a variable of the given type.

```sl
object User {
   name string
}
route POST "/data" {
    from body as user User;
}
```

## Responses:

```sl 
respond with <json|text|html> <value> <status <number>?>
```

Examples:
```sl
respond with { "error": "Something went wrong!" } status 500
respond with status 404
respond with text "Hello world"

object User {
   name string
}
var user User = User{ name: "John" }
respond with user;
```
