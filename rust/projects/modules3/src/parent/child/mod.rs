mod util;

pub mod whine {
    pub fn bath() {
        println!("I don't want to have a bath!")
    }

    pub fn teeth() {
        println!("I did brush my teeth!");
        super::util::noise::cry();
    }
}

