struct StratUnit {
    code: String,
    name: String,
    description: String,
}

impl StratUnit {
    // New returns a new StratUnit with only a code
    fn new(strat_code: &str) -> StratUnit {
        let code = String::from(strat_code);
        StratUnit {
            code: String::from(code),
            name: String::from(""),
            description: String::from(""),
        }
    }

    fn describe(&self) {
        println!("{} ({}) - {}", self.name, self.code, self.description);
    }
}

fn main() {
    let mut tuth = StratUnit::new("tuth");
    tuth.name = String::from("Tutholamia");
    tuth.description = String::from("Tuth is a nice type of rock");
    tuth.describe();
}
