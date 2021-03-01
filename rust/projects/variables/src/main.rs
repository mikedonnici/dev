fn main() {
    let mut x = 5;
    println!("The value of x is: {}", x);
    x = 6;
    println!("The value of x is: {}", x);

    const MAX_NUM: u32 = 9;
    println!("The value of MAX_NUM is: {}", MAX_NUM);

    let _n = 8;
    let _n = 18;
    println!("The value of n is: {}", _n);

    let tup = (500, 6.4, 1);
    let (x, y, z) = tup;
    println!("{}, {}, {}", x, y, z);

    let tup = (500, 6.4, 1);
    println!("{}, {}, {}", tup.0, tup.1, tup.2);
}
