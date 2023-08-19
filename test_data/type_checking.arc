
func testing(val int, pls bool) {
    fmt::printf("testing: val:%v, pls:%v\n", val, pls)
}

func testingg(pls bool) {
}

object AnotherOne {
    x int
    y bool
}
object Yeet {
    yas int
    pls AnotherOne
}

func main(){
    var v = Yeet{
        yas: 1,
        pls: AnotherOne{
            x: 2,
            y: true
        },
    };

    // var pls = 1;
    // var x = ll;

    testingg(v.pls.y)

    var x = 1;

    var vv = v.pls;
    testing(vv.x, vv.y)
    // testing(vv.x, vv.x)

    testing(v.yas, v.pls.y)
    testing(v.yas, v.pls.y)
    testing(x, true)

    // testing(v.pls.y, v.pls.y)
}
