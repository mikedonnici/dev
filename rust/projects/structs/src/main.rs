fn main() {
    struct CountThings {
        foo: u32,
        bar: u32,
        baz: u32,
        bing: u32,
    };

    let things1 = CountThings {
        foo: 8,
        bar: 8,
        baz: 10,
        bing: 10,
    };

    let things2 = CountThings {
        foo: 10,
        bar: 10,
        ..things1
    };

    println!("{}", things2.foo);
    println!("{}", things2.bar);
    println!("{}", things2.baz);
    println!("{}", things2.bing);

}
