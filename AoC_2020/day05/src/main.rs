use std::{fs, time::Instant};

fn main() {
    let content =
        fs::read_to_string("./input/day05.txt").expect("Something went wrong reading the file");
    let passes = parse(&content);

    let time = Instant::now();
    let p1 = part1(&passes);
    let p2 = part2(&passes);
    let duration = time.elapsed();
    println!("Part1: {}", p1);
    println!("Part2: {}", p2);

    println!("Time: {} ns", duration.as_nanos());
}

pub fn part1(ids: &[usize]) -> usize {
    *ids.iter().max().unwrap()
}

pub fn part2(ids: &[usize]) -> usize {
    let min = ids.iter().min().unwrap();
    let max = ids.iter().max().unwrap();

    // XOR everything between the range with all ids
    let all = (*min..=*max).fold(0, |acc, v| acc ^ v);
    all ^ ids.iter().fold(0, |acc, v| acc ^ v)
}

pub fn parse(input: &str) -> Vec<usize> {
    input
        .lines()
        .map(|p| {
            let (row, col) = get_seat(p);
            return row * 8 + col;
        })
        .collect()
}

fn get_seat(p: &str) -> (usize, usize) {
    let mut min_r = 0;
    let mut max_r = 127;
    let mut min_c = 0;
    let mut max_c = 7;

    for c in p.chars() {
        match c {
            'F' => max_r = (max_r + min_r) / 2,
            'B' => min_r = (max_r + min_r) / 2 + 1,
            'L' => max_c = (max_c + min_c) / 2,
            'R' => min_c = (min_c + max_c) / 2 + 1,
            _ => unreachable!(),
        }
    }

    (min_r, min_c)
}

#[cfg(test)]
mod tests {
    use super::*;

    fn example_data() -> Vec<usize> {
        parse(
            "FBFBBFFRLR
BFFFBBFRRR
FFFBBBFRRR",
        )
    }

    #[test]
    fn test_part1() {
        assert_eq!(part1(&example_data()), 567);
    }
}
