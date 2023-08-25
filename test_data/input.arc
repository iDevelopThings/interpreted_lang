object MyObject {
    StrValue string
    StrValue string
    IntValue int
    Nested MyObject
}

func (obj MyObject) TestFunc() int {
    // fmtt::fmt::printf("testing func: %s\n", obj.StrValue)
    var a = 1 + 2 + obj.IntValue;

    return a;
}

func runInputScript() {
    var z MyObject = MyObject{
        StrValue: "root str",
        IntValue: (1 + 2) * 3,
    };

    fmt::printf("testing func: %d\n", z.TestFunc());
    // a.b.thing();

    // var a = 1 + 2;
    // var b = 1 + 2 ^ -3;
    // var z MyObject* = MyObject{
    //     StrValue: "root str",
    //     Nested: MyObject{
    //         StrValue: "nested str",
    //         IntValue: 123,
    //     },
    // };
    // fmt::printf("testing var: %s - nested: %s\n", z.StrValue, z.Nested.StrValue)
}
