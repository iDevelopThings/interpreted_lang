# Arrays

## Declaration

```arc
object User {
   name string
}

var user []User = {
    User{name:"John"},
    User{name:"Bruce"},
};
```

## Slices

```arc
var user []User = {
    User{name:"John"},
    User{name:"Bruce"},
};

// Start from 0, end at 1
var a = user[0:1];

// Start from 1, end at 2
var b = user[1:2];

// Start from 0, end at (len(user))
var c = user[0:];

// Get element at index 0
var d = user[0];
```

## Api

```arc
var u = User{name:"Barry"};

// Add element
users.add(u);

// Remove via object reference
users.remove(u);

// Remove via index
users.remove(0);

delete users[0];

// Get length
users.length();

```