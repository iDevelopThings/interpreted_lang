# Objects

## Declaration

```sl
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

```sl
func (o MyObject) Print() {
    fmt::printf("StrValue: %s\n", o.StrValue);
    fmt::printf("IntValue: %d\n", o.IntValue);
}
```

## Accessing fields

```sl
var obj MyObject = MyObject{
    StrValue: "root str",
    IntValue: (1 + 2) * 3,
};

var str string = obj.StrValue;
```