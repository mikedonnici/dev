

fn main() {
    let n: u8 = 43;
    match n {
        1 => {
            println!("{}", "One");
        }
        2 => {
            println!("{}", "Two");
        }
        3 => {
            println!("{}", "Three");
        }
        _ => () // unit value - nada
    }
}
