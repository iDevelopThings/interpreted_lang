
func main() {
  defer {
    fmt::printf("testing\n");
  }

  defer func (param type) {
    fmt::printf("testing\n");
  }
}