# Enums

**Need more work, but a basic implementation exists/works**

### Declaration

```arc
enum Color {
    Red = "red"
    Blue = 1 // because why not
    Green = true // again?
    SomeValueHoldingRuntimeValues(a int, b string)
}
```

### Usage

```arc
var red Color = Color.Red;
var red Color = Color::Red;

var pls Color = Color::SomeValueHoldingRuntimeValues(1, "hello");
pls.a; // 1
pls.b; // "hello"

var eq = pls == Color::SomeValueHoldingRuntimeValues

var a = red == Color.Red; // true
```


