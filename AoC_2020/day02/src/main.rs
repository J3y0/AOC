use std::fs;
use std::time::Instant;

pub struct Policy<'a> {
    a: usize,
    b: usize,
    letter: u8,
    password: &'a [u8],
}

impl Policy<'_> {
    fn from([a, b, c, d]: [&str; 4]) -> Policy<'_> {
        Policy {
            a: a.parse().unwrap(),
            b: b.parse().unwrap(),
            letter: c.as_bytes()[0],
            password: d.as_bytes(),
        }
    }
}

fn main() {
    let content =
        fs::read_to_string("./input/day02.txt").expect("Something went wrong reading the file");
    let passwords = parse(&content);

    let time = Instant::now();
    let p1 = part1(&passwords);
    let p2 = part2(&passwords);
    let duration = time.elapsed();

    println!("Part1 : {}", p1);
    println!("Part2 : {}", p2);
    println!("Time: {} ns", duration.as_nanos());
}

pub fn parse(input: &str) -> Vec<Policy<'_>> {
    let temp: Vec<&str> = input
        .split(['-', ':', ' ', '\n'])
        .filter(|s| !s.is_empty()).collect();

    temp.chunks(4).map(|c| Policy::from([c[0], c[1], c[2], c[3]])).collect()
}

pub fn part1(lines: &[Policy<'_>]) -> usize {
    lines
        .iter()
        .filter(|&p| {
            let count = p.password.iter().filter(|&c| *c == p.letter).count();
            count >= p.a && count <= p.b
        })
        .count()
}

pub fn part2(lines: &[Policy<'_>]) -> usize {
    lines
        .iter()
        .filter(|&p| {
            let first = p.password[p.a - 1] == p.letter;
            let second = p.password[p.b - 1] == p.letter;
            first ^ second
        })
        .count()
}

#[cfg(test)]
mod tests {
    use super::*;

    fn example_data()-> Vec<Policy<'static>> {
        parse("1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc")
    }

    #[test]
    fn part1_test() {
        assert_eq!(part1(&example_data()), 2);
    }

    #[test]
    fn part2_test() {
        assert_eq!(part2(&example_data()), 1);
    }
}