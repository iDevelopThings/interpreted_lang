
func pls() {
	var value = {
        "key1": "value1",
        "key2": "value2",
    };

    var key1 = "key1";
    var key2 = "key2";

    fmt::printf("Value: %v\n", value);
    fmt::printf("Key1: %v\n", value.key1);
    fmt::printf("Key2: %v\n", value.key2);
    fmt::printf("Key1: %v\n", value["key1"]);
    fmt::printf("Key2: %v\n", value["key2"]);

    var key1 = "key1";
    var key2 = "key2";
    fmt::printf("Key1: %v\n", value[key1]);
    fmt::printf("Key2: %v\n", value[key2]);

    value.key1 = "value1.1";
    value.key2 = "value2.1";

    value["key3"] = "value3";
    fmt::printf("Value: %v\n", value);

    value["key3"] = "value3.1";
    fmt::printf("Value: %v\n", value);

    var key4 = "key4";
    value[key4] = "value4";

    value.key5 = "value5";

    fmt::printf("Value: %v\n", value);
}

func main() {

    pls()
    pls()
    pls()
    pls()
    pls()
    pls()

}