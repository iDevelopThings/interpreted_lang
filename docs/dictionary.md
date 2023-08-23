# Dictionary/Map

## Declaration

```arc
var dict = {
    "key1": "value1",
    "key2": 2,
};
```

## Get value

```arc
// Via string key
dict["key1"];

// Via variable key
var k = "key2";
dict[k];

// Direct member access
dict.key1;
```

## Set value

```arc
dict["key"] = value;
var k = "key";
dict[k] = value;
dict.key = value;

delete dict.key;

```

