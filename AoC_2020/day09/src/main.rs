use std::{fs, time::Instant};

fn main() {
    let content =
        fs::read_to_string("./input/day09.txt").expect("Something went wrong reading the file");

    let numbers = parse(&content);

    let time = Instant::now();
    let p1 = part1(&numbers, 25).unwrap();
    let p2 = part2(&numbers, &p1);
    let duration = time.elapsed();
    println!("Part1: {}", p1);
    println!("Part2: {}", p2);

    println!("Time: {} ns", duration.as_nanos());
}

pub fn part1(numbers: &[u64], window_size: usize) -> Option<u64> {
    numbers
        .windows(window_size + 1)
        .find(|w| {
            for i in 0..window_size + 1 {
                for j in i + 1..window_size + 1 {
                    if w[i] + w[j] == w[window_size] {
                        return false;
                    }
                }
            }
            true
        })
        .map(|w| w[window_size])
}

pub fn part2(numbers: &[u64], invalid: &u64) -> u64 {
    let mut start = 0;
    let mut end = 2;
    let mut sum = numbers[0] + numbers[1];

    while sum != *invalid {
        if sum < *invalid {
            sum += numbers[end];
            end += 1;
        } else {
            sum -= numbers[start];
            start += 1;
        }
    }

    let w = &numbers[start..end];
    let min = w.iter().min().unwrap();
    let max = w.iter().max().unwrap();

    min + max
}

fn parse(input: &str) -> Vec<u64> {
    input.lines().map(|s| s.parse().unwrap()).collect()
}

#[cfg(test)]
mod tests {
    use super::*;

    fn example_data() -> Vec<u64> {
        parse(
            "35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576",
        )
    }

    #[test]
    fn test_part1() {
        assert_eq!(part1(&example_data(), 5), Some(127));
    }

    #[test]
    fn test_part2() {
        assert_eq!(part2(&example_data(), &127), 62);
    }
}
