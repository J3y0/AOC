use std::{fs, time::Instant};

fn main() {
    let content =
        fs::read_to_string("./input/day01.txt").expect("Something went wrong reading the file");
    let numbers = parse(&content);

    let now = Instant::now();
    let p1 = part1(&numbers).unwrap();
    let p2 = part2(&numbers).unwrap();
    let time = now.elapsed();

    println!("Part 1: {}", p1);
    println!("Part 2: {}", p2);
    println!("Time: {} ns", time.as_nanos());
}

pub fn parse(input: &str) -> Vec<usize> {
    input.lines().map(|x| x.parse().unwrap()).collect()
}

pub fn part1(input: &[usize]) -> Option<usize> {
    let mut hash = [0; 2020];
    sum_two(input, &mut hash, 2020, 1)
}

pub fn part2(input: &[usize]) -> Option<usize> {
    for i in 0..input.len() {
        let n = input[i];
        let mut hash = [0; 2020];
        if let Some(res) = sum_two(&input[i + 1..], &mut hash, 2020 - n, i + 1) {
            return Some(n * res);
        }
    }
    None
}

fn sum_two(slice: &[usize], hash: &mut [usize], target: usize, val: usize) -> Option<usize> {
    for &n in slice {
        if n > target {
            continue;
        }
        if hash[n] == val {
            return Some(n * (target - n));
        }
        hash[target - n] = val;
    }
    None
}

#[cfg(test)]
mod tests {
    use super::*;

    fn example_data() -> Vec<usize> {
        parse(
            "1721
979
366
299
675
1456",
        )
    }

    #[test]
    fn part1_test() {
        assert_eq!(part1(&example_data()), Some(514579));
    }

    #[test]
    fn part2_test() {
        assert_eq!(part2(&example_data()), Some(241861950));
    }
}
