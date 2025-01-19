use std::{fs, time::Instant};

fn main() {
    let content =
        fs::read_to_string("./input/day10.txt").expect("Something went wrong reading the file");

    let adapters = parse(&content);

    let time = Instant::now();
    let p1 = part1(&adapters);
    let p2 = part2(&adapters);
    let duration = time.elapsed();
    println!("Part1: {}", p1);
    println!("Part2: {}", p2);

    println!("Time: {} ns", duration.as_nanos());
}

pub fn part1(adapters: &[usize]) -> usize {
    let mut delta = [0, 0, 0, 1];
    delta[adapters[0]] += 1;

    adapters.windows(2).for_each(|w| delta[w[1] - w[0]] += 1);

    delta[1] * delta[3]
}

pub fn part2(adapters: &[usize]) -> usize {
    let max = *adapters.iter().last().unwrap();
    let mut mem = vec![0; max + 1];

    mem[0] = 1;
    for &a in adapters {
        match a {
            1 => mem[a] = mem[a - 1],
            2 => mem[a] = mem[a - 1] + mem[a - 2],
            _ => mem[a] = mem[a - 1] + mem[a - 2] + mem[a - 3],
        }
    }

    mem[max]
}

fn parse(input: &str) -> Vec<usize> {
    let mut outlets: Vec<_> = input.lines().map(|s| s.parse().unwrap()).collect();
    outlets.sort();

    outlets
}

#[cfg(test)]
mod tests {
    use super::*;

    fn example_data() -> Vec<usize> {
        parse(
            "16
10
15
5
1
11
7
19
6
12
4",
        )
    }

    #[test]
    fn test_part1() {
        assert_eq!(part1(&example_data()), 35);
    }

    #[test]
    fn test_part2() {
        assert_eq!(part2(&example_data()), 8);
    }
}
