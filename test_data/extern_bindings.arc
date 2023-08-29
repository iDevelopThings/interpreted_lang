object http {
}

object RequestOptions {
  data dict
}

extern func (http) get(url string) string

func main() {

    var response string = http::get("http://pokeapi.co/api/v2/pokedex/kanto/")
    fmt::printf("Response: %s\n", response)

}