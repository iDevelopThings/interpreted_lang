func testOption() ?string {
  return none
}

func main() {
  if testOption() == none {
	fmt::print("...")
  }


  //  var opt = testOption()
  //  var res = testResult() or {
  //	fmt::print("...")
  //	return "hi"
  //  }

  //  var anothaOne = testResult() or fmt::print("...")
}

//func testResult() !string {}
