
object User {
   name string
}
func (u User) String() string {
    return u.name;
}

func pls() {
	var users []User = {
	    User{name:"John"},
	    User{name:"Bruce"},
	};


	fmt::printf("Len: %s\n", users.length());

    var barry = User{name:"Barry"};
    users.add(barry);

	fmt::printf("Len: %s\n", users.length());

    users.remove(0);
	users.remove(barry);

	fmt::printf("Len: %s\n", users.length());

  	for users {
		fmt::printf("User: %v\n", it);
	}

  	for users as u {
		fmt::printf("User(u): %v\n", u);
	}
}

func main() {

    pls()

}