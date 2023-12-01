use std::fs;

fn main() {
    let content: String = fs::read_to_string("./data/day4.txt")
        .expect("Should have been able to read the file");
    
    let lines: Vec<&str> = content.lines().collect();
    let mut part1: usize = 0;
    let mut part2: usize = 0;

    for line in lines {
        match parse(line) {
            Ok(pairs) => {
                if pair_contains_the_other((pairs[0], pairs[1]), (pairs[2], pairs[3])) {
                    part1 += 1;
                    part2 += 1;
                } else if overlap((pairs[0], pairs[1]), (pairs[2], pairs[3])) {
                    part2 += 1;
                }
            },
            Err(_) => println!("Error while parsing {line}")
        }
    }

    println!("Part1: {part1}");
    println!("Part2: {part2}");
}

fn parse(line: &str) -> Result<[u8; 4], std::num::ParseIntError> {
    let values: Vec<Vec<u8>> = line.split(",")
        .map(|item| item
                .split("-")
                .map(|nb| match nb.parse() {Ok(num) => num, Err(_) => 0})
                .collect()
            )
        .collect();

    let mut result: [u8; 4] = [0; 4];
    result[0] = values[0][0];
    result[1] = values[0][1];
    result[2] = values[1][0];
    result[3] = values[1][1];
    Ok(result)
}

// Part 1
fn within(pair1: (u8, u8), pair2: (u8, u8)) -> bool {
    return ((pair2.0 <= pair1.0) && (pair1.0 <= pair2.1))
        && ((pair2.0 <= pair1.1) && (pair1.1 <= pair2.1));
}

fn pair_contains_the_other(pair1: (u8, u8), pair2: (u8, u8)) -> bool {
    return within(pair1, pair2) || within(pair2, pair1);
}

// Part 2
fn overlap(pair1: (u8, u8), pair2: (u8, u8)) -> bool {
    return pair1.0 <= pair2.1 && pair1.1 >= pair2.0;
}