use std::collections::BTreeMap;
use std::fs;

fn main() {
    let content =
        fs::read_to_string("./data/day1.txt").expect("Should have been able to read the file");

    let lines: Vec<&str> = content.lines().collect();
    let calories: BTreeMap<i32, i16> = parse(lines);

    // Part 1
    let part1 = calories.iter().next_back().unwrap();
    println!("{:?}", part1);

    // Part 2
    let mut sum: i32 = 0;
    let keys: Vec<&i32> = calories.keys().collect();
    for i in 0..3 {
        sum += keys[keys.len() - i - 1];
    }
    println!("The sum is: {sum}");
}

fn parse(lines: Vec<&str>) -> BTreeMap<i32, i16> {
    let mut calories: BTreeMap<i32, i16> = BTreeMap::new();
    let mut nb_elf: i16 = 0;
    let mut sum: i32 = 0;

    for i in 1..lines.len() {
        let value: i32 = match lines[i].parse() {
            Ok(num) => num,
            Err(_) => -1, // The case where there is an empty string as a line
        };

        if value == -1 {
            calories.insert(sum, nb_elf);
            nb_elf += 1;
            sum = 0;
        } else {
            sum += value;
        }
    }
    return calories;
}
