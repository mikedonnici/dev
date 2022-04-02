mod some_mod_name {

    pub struct Foo {
        pub a: String,
        b: String,
    }

    impl Foo {
        pub fn new(a_val: String) -> Foo {
            Foo {
                a: a_val,
                b: String::from("default b"),
            }
        }
    }
}
