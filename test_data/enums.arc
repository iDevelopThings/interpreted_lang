enum Yeet {
  a = 1
  b = 2
  WithValue(int, Obj)
}

object Obj {
  name string
}

func (o Obj) String() string {
  return "...";
}

enum Yeet2 {
  Another(val string)
  WithValue(a int, b Obj)
}

func main() {
  var a Yeet = Yeet::WithValue(1, Obj{name: "a"});
  //  fmt::printf("WithValue -> a: %v\n", Yeet::a);
  //  fmt::printf("WithValue -> b: %v\n", Yeet::b);
  //  var idx int = 0;
  //  fmt::printf("WithValue -> 0: %v\n", a[idx]);
  //  fmt::printf("WithValue -> 1: %v\n", a[1].name);

  if a["1"] == Yeet::a {
	fmt::println("a == Yeet::a");
  } else {
	fmt::println("a != Yeet::a");
  }

}