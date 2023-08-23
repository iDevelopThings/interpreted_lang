## Throw a panic

```arc
func main() {
  error::panic("Something went wrong")
}
```

## Types

Option types: `?` mean the value could be `none`(null | nil)

```arc
func getUser() ?User {}
```

Result types: `!` mean the value could contain/hold an `error`

```arc
func getUser() !User {}
```

## Error handling

```arc
func main() {
  // getUser = return type: !User
  var user = getUser() or {
	// handle error
	error::panic("Something went wrong")
  }
  var user = getUser() or fmt::println("Something went wrong")
}
```

## Option handling

```arc
func main() {
  // getUser = return type: ?User
  var user = getUser() or User{name:"John"}
}
```