pub fn is_leap_year(year: u64) -> bool {
    let y = &year;
    (y % 4 == 0 && !(y % 100 == 0)) || (y % 4 == 0 && y % 400 == 0)
}
