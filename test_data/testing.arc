object MyTestingObject {
  name string
  age  int
}


func (ob MyTestingObject) String() string {
  //  var x MyTestingObject;
  //  x.name;
  ob.name = ".."
  ob.age = ".."
  ob.pls = ""

  ob.fakyeahh();

  return ob.String();

}

func (ob MyTestingObject) fakyeahh() {

}


func main() {}