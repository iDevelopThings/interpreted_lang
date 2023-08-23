func testOption(i int) ?string {
  if i == 0 {
	return "hi"
  }
  return none
}

func main() {
  var i int = 0
  if testOption(i) == "hi" && testOption(i) == none {
	fmt::println("we is hi")
  }
  i++
  if testOption(i) == none {
	fmt::println("we is none")
  }
}










//func main() {
//  subFunc();
//}
//
//func subFunc() {
//  defer func() {
//	fmt::println("hi from subfunc defer");
//  }
//  anotherSubFunc();
//}
//
//func anotherSubFunc()  {
//  defer {
//	fmt::println("hi from defer");
//  }
//  //  error::panic("hi sir");
//}
