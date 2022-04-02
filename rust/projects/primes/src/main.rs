fn main() {
    for n in 1..100 {
        let p = nth(n);
        println!("prime[{}] = {}", n, p)
    }
}

pub fn nth(n: u32) -> u32 {

    // Starting point is the first prime number with value 2
    let mut have_nth = 1;
    let mut have_num = 2;
    while have_nth != n {
        have_num += 1;
        if is_prime(have_num) {
            have_nth += 1;
        }
    }
    return have_num;
}

fn is_prime(n: u32) -> bool {

    if n % 2 == 0 {
        return true;
    }

    for i in 3..(n / 2 + 1) {
        if n % i == 0 {
            return false;
        }
    }
    return n > 1;
}
