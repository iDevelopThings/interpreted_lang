func testOption(i int) ?string {
  if i == 0 {
	return "hi"
  }
  return none
}

func main() {
  var i int = 0
  if testOption(1) == none {
	fmt::println("3. we is hi")
  }
//  if testOption(0) == "hi" {
//	fmt::println("1. we is hi")
//  }
//  if testOption(1) == "hi" {
//	fmt::println("2. we is hi")
//  }
//  if testOption(i) == "hi" && testOption(i) == none {
//	fmt::println("we is hi")
//  }
  //  i++
  //  if testOption(i) == none {
  //	fmt::println("we is none")
  //  }
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
