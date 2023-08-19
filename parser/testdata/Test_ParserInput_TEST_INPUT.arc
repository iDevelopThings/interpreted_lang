import "...";
import "..."

func NoReceiverFunc() {

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

    var i int = 0;
    for {
        fmt::printf("Looping: %d\n", i)
        if i == 10 {
            break;
        }
        i++;
    }

    if value["key1"] == "value1" {
        return "Hello World"
    } else if 1 == 2 {
        return "nah"
    } else {
        return "Goodbye World"
    }

	var value = {
        "key1": "value1",
        "key2": "value2",
    };




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

}

object MyObject {
    StrValue string
    IntValue int
    Nested MyObject
}

func (a Type) ReceiverFunc() string {
    return "Hello World"
}

func (a MyObject) VoidReceiverFunc() {
    a.StrValue = "Hello World"
}
