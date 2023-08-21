enum Yeet {
  a = 1
  b = 2
  WithValue(int, Obj)
}

object Obj {
  name string
}

enum Yeet2 {
  Another(val string)
  WithValue(a int, b Obj)
}

func main() {
  //  var a Yeet = Yeet::WithValue(1, Obj{name: "a"});

}