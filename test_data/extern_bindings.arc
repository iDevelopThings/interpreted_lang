object http {
}

object RequestOptions {
  data dict
}

extern func (http) get(url string, options dict) dict
extern func (http) post(url string, options dict) dict

func main() {

    // var response dict = http::get("https://pls.idt.link/gib_data.php", {
    //     "body": {"message" : "hi"},
    //     "headers" : {"content-type": "application/json"}
    // })
    var response dict = http::post("https://pls.idt.link/gib_data.php", {
        "body": {"message" : "hi"},
        "headers" : {"content-type": "application/json"}
    })

    fmt::printf("Response: %s\n", response)
    fmt::printf("Response: %s\n", response.message)

}