## Variables

```sl
var <name> <type> = <value>;
var <name> <type>;
```

```sl
var name string = "John";
var name string;
```

## If statements

```sl
if <condition> {
    <body>
} else if <condition> {
    <body>
} else {
    <body>
}
```

```sl
if 1 == 1 {
    print("1 is equal to 1");
}
```

## For loops

```sl

// Regular `while` loop, you control the flow

var i int = 0;
for {
    fmt::printf("Looping: %d\n", i)
    if i == 10 {	
        break;
    }
    i++;
}

for 0..10 {
    // `it` is bound to the loop scope
    fmt::printf("[loop - it] %d\n", it);
}

// Define a custom variable name, instead of `it`
for 0..10 as i {
    fmt::printf("[loop - i] %d\n", i);
}

// Define a custom step, increments `it/i` by n each iteration
for 1..10 step 2 {
    fmt::printf("[loop - it] %d\n", it);
}

// Easily looping over arrays

var users = {
    User{name: "John"},
    User{name: "Jane"},
    User{name: "Bob"},
};

// Via var reference, using `it`
for users {
    fmt::printf("[loop - it] %s\n", it.name);
}

// Via var reference, using a custom name
for users as u {
    fmt::printf("[loop - u] %s\n", u.name);
}

```

## Deletion

Using C++/PHP like syntax for deletion

### Dictionaries
```sl
var dict = {"key": "value"};
delete dict["key"];
delete dict.key;
delete dict[someVarWithKeyName];
```

### Arrays
```sl
var arr = [1, 2, 3];
delete arr[0];
delete arr[someVarWithIndex];

// Using slices
// Deletes the first 2 elements
delete arr[0:2];
// Delete from 1 to the end
delete arr[1:];

```

### Variables
This mainly just deletes the variable from the current scope
Any further access to the variable will result in an error/panic
```sl
var name = "John";

delete name;
```