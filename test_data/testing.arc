object fmt {}

func (fmt) printf(string, ...) {}







object User {
   name string
   id int
   acc Account
}



  func(u User) Something() {
    return "";
  }
func (u User) String() string {
//    return u.acc.fkdjfdkljdkl
  return u.pls;
    return u.id;
  //test

}



func testing(a int, b string) {

}

object User {
   name string
   id int
   f float
   int float
}



func (u User) String() string {
    return u.name;
}

func (User) MyStaticFunc() string {
    return "static";
}


func main() {
    fmt::print(1, true, 1.0f, 1.0f, "test", "\n");
    fmt::print("test\n");
    fmt::println("test", 1);

    var user = User{
        name: "John",
        id: 0,
        f: 0.0f
    };

    fmt::printf("obj: %v\n", user);

    user.String();
    // testing(false, 1);

    fmt::printf("user.String() = %s\n", user.String());
    fmt::printf("User::MyStaticFunc() = %s\n", User::MyStaticFunc());

}
