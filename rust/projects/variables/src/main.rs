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

    let a1 = [1, 2, 3, 4, 5];
    println!("{}", a1[3]);

    let first_num = 2;
    let second_num = 4;
    let s = sum_numbers(2, 4);
    println!("Sum of {} and {} is {}", first_num, second_num, s);

    let n = 4;
    if n > 5 {
        println!("n is greater than 5");
    } else if n == 4 {
        println!("n is equal to 4");
    } else {
        println!("n is less than 4");
    }

    let condition = true;
    let number = if condition { 5 } else { 6 };
    println!("The value of number is: {}", number);

    let mut count = 0;
    let result = loop {
        count += 1;
        println!("Count is now {}", count);
        if count >= 4 {
            println!("break and return value from loop");
            break count + 100;
        }
    };
    println!("result from loop is {}", result);

    let a = [1, 2, 3];
    for n in a.iter() {
        println!("{}", n);
    }

    for n in 1..4 {
        println!("{}", n);
    }

    for n in (1..4).rev() {
        println!("{}", n);
    }

    let s1 = String::from("Mike");
    let s2 = s1.clone();
    println!("s1 = {}", s1);
    println!("s2 = {}", s2);

    let a = 5;
    let b = a;
    println!("a and b are still both in scope, a = {}, b = {}", a, b);

    // Here a and b stay in scope as they are stack values
    let a = 5;
    let b = 7;
    let sum = sum_numbers(a, b);
    println!("{}, {}", a, b);

    // But here they are deactivated (or moved)
    let a = String::from("dog");
    let b = String::from("cat");
    print_strings(a, b);
    // println!("{}, {}", a, b); // <- compilation error

    let mut s = String::from("abc");
    let len = str_len(&s);
    println!("string_len(\"{}\") = {}", s, len);

    change_str(&mut s)


}

fn sum_numbers(a: u32, b: u32) -> u32 {
    return a + b;
}

fn print_strings(a: String, b: String) {
    println!(
        "Function now owns String args {} and {} - they are 'moved'",
        a, b
    );
}

fn str_len(s: &String) -> usize {
    s.len()
}

fn change_str(s: &mut String) {
   s.push_str(", world");
}
