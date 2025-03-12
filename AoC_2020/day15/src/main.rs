use std::{fs, time::Instant};

fn main() {
    let content =
        fs::read_to_string("./input/day15.txt").expect("Something went wrong reading the file");

    let numbers = parse(&content);

    let time = Instant::now();
    let p1 = part1(&numbers);
    let p2 = part2(&numbers);
    let duration = time.elapsed();
    println!("Part1: {}", p1);
    println!("Part2: {}", p2);

    println!("Time: {} ns", duration.as_nanos());
}

pub fn part1(input: &[usize]) -> usize {
    play(input, 2020)
}

pub fn part2(input: &[usize]) -> usize {
    play(input, 30_000_000)
}

fn play(input: &[usize], limit: usize) -> usize {
    let size = input.len();
    let mut numbers = vec![0; limit];

    for i in 0..size {
        numbers[input[i]] = i + 1;
    }

    let mut cur = input[size - 1];
    for i in size..limit {
        let prev_turn = numbers[cur];
        numbers[cur] = i;
        cur = if prev_turn == 0 { 0 } else { i - prev_turn };
    }

    cur
}

fn parse(input: &str) -> Vec<usize> {
    input.split(',').map(|n| n.parse().unwrap()).collect()
}

#[cfg(test)]
mod tests {
    use super::*;

    fn example_data() -> Vec<usize> {
        parse("0,3,6")
    }

    #[test]
    fn test_part1() {
        assert_eq!(part1(&example_data()), 436);
    }

    #[test]
    fn test_part2() {
        assert_eq!(part2(&example_data()), 175594);
    }
}
