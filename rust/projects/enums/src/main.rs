#[derive(Debug)]
enum Fruit {
    Apple,
    Orange,
    Banana,
}

impl Fruit {
    fn describe(&self) {
        println!("Fruit is {:?}", self);
    }
}

fn main() {
    let f1 = Fruit::Apple;
    let f2 = Fruit::Banana;
    f1.describe();
    f2.describe();
}
