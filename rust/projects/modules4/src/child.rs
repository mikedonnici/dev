mod util; // <- import util.rs into 'util' namespace

pub mod whine {
    pub fn bath() {
        println!("I don't want to have a bath!")
    }

    pub fn teeth() {
        println!("I did brush my teeth!");
        crate::util::noise::cry();
    }
}

