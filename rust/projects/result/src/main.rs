fn main() {
    let res = do_thing(true);
    match res {
        Ok(c) => println!("Result ok, code = {}", c),
        Err(e) => println!("Error, code = {}", e),
    }
}

fn do_thing(ret_err: bool) -> Result<i32,i32> {
    if ret_err {
        return Err(1);
    }
    Ok(0)
}

