# Objects

## Declaration

```arc
object MyObject {
    StrValue string
    IntValue int
    Nested MyObject
}

var obj MyObject = MyObject{
    StrValue: "root str",
    IntValue: (1 + 2) * 3,
};
```


## Methods

### Public
```arc
func (o MyObject) Print() {
    fmt::printf("StrValue: %s\n", o.StrValue);
    fmt::printf("IntValue: %d\n", o.IntValue);
}
```

### Static
The missing receiver var indicates a static method.

The functions are bound to the type and not to an instance of the type.
```arc
func (MyObject) Print() {
    
}
```

## Accessing fields

```arc
var obj MyObject = MyObject{
    StrValue: "root str",
    IntValue: (1 + 2) * 3,
};

var str string = obj.StrValue;
```