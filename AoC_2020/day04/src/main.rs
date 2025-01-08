use std::{fs, time::Instant};

type Passport<'a> = Vec<[&'a str; 2]>;

fn main() {
    let content =
        fs::read_to_string("./input/day04.txt").expect("Something went wrong reading the file");
    let passports = parse(&content);

    let time = Instant::now();
    let p1 = part1(&passports);
    let p2 = part2(&passports);
    let duration = time.elapsed();
    println!("Part1: {}", p1);
    println!("Part2: {}", p2);

    println!("Time: {} ns", duration.as_nanos());
}

pub fn part1(passports: &[Passport]) -> usize {
    passports.iter().filter(|&p| p.len() == 7).count()
}

pub fn part2(passports: &[Passport]) -> usize {
    passports
        .iter()
        .filter(|&p| p.len() == 7)
        .filter(|&p| p.iter().all(validate))
        .count()
}

pub fn parse(input: &str) -> Vec<Passport> {
    input.split("\n\n").map(parse_block).collect()
}

fn parse_block(block: &str) -> Passport {
    let mut fields = Vec::with_capacity(7);

    let splitted: Vec<_> = block
        .split([' ', ':', '\n'])
        .filter(|s| !s.is_empty())
        .collect();
    for i in (0..splitted.len()).step_by(2) {
        if splitted[i] == "cid" {
            continue;
        }
        fields.push([splitted[i], splitted[i + 1]]);
    }

    fields
}

fn validate(field: &[&str; 2]) -> bool {
    let value = field[1];
    match field[0] {
        "byr" => value.parse().is_ok_and(|v| 1920 <= v && v <= 2002),
        "iyr" => value.parse().is_ok_and(|v| 2010 <= v && v <= 2020),
        "eyr" => value.parse().is_ok_and(|v| 2020 <= v && v <= 2030),
        "hgt" => {
            if value.ends_with("in") {
                return value[..2].parse().is_ok_and(|v| 59 <= v && v <= 76);
            } else if value.ends_with("cm") {
                return value[..3].parse().is_ok_and(|v| 150 <= v && v <= 193);
            } else {
                return false;
            }
        }
        "hcl" => {
            value.starts_with('#')
                && value.len() == 7
                && value[1..].bytes().all(|b| b.is_ascii_hexdigit())
        }
        "ecl" => ["amb", "blu", "brn", "gry", "grn", "hzl", "oth"].contains(&value),
        "pid" => value.len() == 9 && value.bytes().all(|b| b.is_ascii_digit()),
        _ => unreachable!(),
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    fn example_data() -> Vec<Passport<'static>> {
        parse(
            "ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929

hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm

hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in",
        )
    }

    #[test]
    fn test_part1() {
        assert_eq!(part1(&example_data()), 2);
    }
}
