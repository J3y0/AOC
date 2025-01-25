use std::{fs, time::Instant};

fn main() {
    let content =
        fs::read_to_string("./input/day06.txt").expect("Something went wrong reading the file");

    let answers = parse(&content);

    let time = Instant::now();
    let p1 = part1(&answers);
    let p2 = part2(&answers);
    let duration = time.elapsed();
    println!("Part1: {}", p1);
    println!("Part2: {}", p2);

    println!("Time: {} ns", duration.as_nanos());
}

pub fn part1(answers: &[u32]) -> u32 {
    let mut sum = 0;
    let mut group: u32 = 0;
    for &p in answers {
        if p == 0 {
            sum += group.count_ones();
            group = 0;
        } else {
            group |= p;
        }
    }
    sum += group.count_ones();

    sum
}

pub fn part2(answers: &[u32]) -> u32 {
    let mut sum = 0;
    let mut group = u32::MAX;
    for &p in answers {
        if p == 0 {
            sum += group.count_ones();
            group = u32::MAX;
        } else {
            group &= p;
        }
    }
    sum += group.count_ones();

    sum
}

fn parse(input: &str) -> Vec<u32> {
    input
        .lines()
        .map(|l| l.bytes().fold(0, |acc, b| acc | (1 << (b - b'a'))))
        .collect()
}

#[cfg(test)]
mod tests {
    use super::*;

    fn example_data() -> Vec<u32> {
        parse(
            "abc

a
b
c

ab
ac

a
a
a
a

b",
        )
    }

    #[test]
    fn test_part1() {
        assert_eq!(part1(&example_data()), 11);
    }

    #[test]
    fn test_part2() {
        assert_eq!(part2(&example_data()), 6);
    }
}
