
// import "imports/nope.arc"

object MyTestingObject {
  name string
  age  int
  pls  string
}

func testOption() ?string {
  fmt::printf("");
  return none
}

func (ob MyTestingObject) String() string {
  return "hi..."
}

func (ob MyTestingObject) TestSettingData(testing string) {
  ob.name = ".."
  //    ob.age = ".."
  ob.pls = ""
}
func (ob MyTestingObject) Yayeeet(pls string, string bool) {
}


func main() {


  var obj MyTestingObject = MyTestingObject{
	name : "Bruce"
  }

  obj.Yayeeet("yayay", false);

}