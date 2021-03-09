fn main() {
    let s = String::from("Hello, world!");
    let fw = first_word(&s);
    println!("First word: {}", fw)
}

fn first_word(s: &str) -> &str {
    let bytes = s.as_bytes();
    for (i, &item) in bytes.iter().enumerate() {
        if item == b' ' {
            return &s[0..i];
        }
    }
    &s[..]
}
