
pub mod parent {

    pub mod help {
        pub fn homework() {
            println!("Yes, I can help with your homework");
        }
        pub fn playdate() {
            println!("Yes, I can drive you to your play date");
        }
    }

    pub mod hinder {
        pub fn bedtime() {
            println!("Go to bed!");
        }
        pub fn clean_room() {
            println!("Clean your room!");
        }
    }
}

pub mod child {
    pub mod whine {
        pub fn bath() {
            println!("I don't want to have a bath!")
        }
        pub fn teeth() {
            println!("I did brush my teeth!")
        }
    }
}
