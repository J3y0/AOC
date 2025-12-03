use std::collections::BTreeSet;

use crate::days::Solution;

pub struct Range {
    start: String,
    end: String,
    start_parsed: usize,
    end_parsed: usize,
}

impl Range {
    fn new(start: &str, end: &str) -> Self {
        let s_parsed = start.parse().unwrap();
        let e_parsed = end.parse().unwrap();
        Range {
            start: String::from(start),
            end: String::from(end),
            start_parsed: s_parsed,
            end_parsed: e_parsed,
        }
    }
}

pub struct Day02;

impl Solution for Day02 {
    type Input = Vec<Range>;

    fn parse(data: &str) -> Self::Input {
        data.trim_end()
            .split(',')
            .map(|r| {
                let (start, end) = r.split_once('-').unwrap();
                Range::new(start, end)
            })
            .collect()
    }

    fn part1(input: &Self::Input) -> usize {
        let mut tot = 0;
        for range in input {
            let half_start: usize = range.start[..range.start.len() / 2].parse().unwrap_or(1);
            let mut half_end: usize = range.end[..range.end.len() / 2].parse().unwrap_or(1);

            if half_end < half_start {
                half_end *= 10;
            }

            tot += (half_start..=half_end)
                .map(|e| e * 10_usize.pow(e.ilog10() + 1) + e) // concat x||x
                .filter(|e| (range.start_parsed..=range.end_parsed).contains(e))
                .sum::<usize>();
        }

        tot
    }

    fn part2(input: &Self::Input) -> usize {
        let mut tot = 0;
        for range in input {
            let mut seen: BTreeSet<usize> = BTreeSet::new();
            let max_digit_count = range.end.len() / 2;
            for nb_digit in 1..=max_digit_count {
                let pat = range.start[..nb_digit].parse().unwrap_or(1);
                let mut pat_end = range.end[..nb_digit].parse().unwrap_or(1);

                if pat_end < pat {
                    // To account cases such as 9 - 23
                    for val in 0..=pat_end {
                        let repeated = concat(val, range.end.len() / nb_digit);
                        if (range.start_parsed..=range.end_parsed).contains(&repeated) {
                            if seen.insert(repeated) {
                                tot += repeated
                            }
                        }
                    }
                    pat_end *= 10;
                }

                for val in pat..=pat_end {
                    // To account cases where nb_digits start != nb_digits end
                    let repeated = concat(val, range.end.len() / nb_digit);
                    if (range.start_parsed..=range.end_parsed).contains(&repeated) {
                        if seen.insert(repeated) {
                            tot += repeated
                        }
                    }

                    if nb_digit == range.start.len() {
                        continue;
                    }
                    let repeated = concat(val, range.start.len() / nb_digit);
                    if (range.start_parsed..=range.end_parsed).contains(&repeated) {
                        if seen.insert(repeated) {
                            tot += repeated
                        }
                    }
                }
            }
        }

        tot
    }
}

fn concat(val: usize, times: usize) -> usize {
    val.to_string().repeat(times).parse().unwrap()
}

#[cfg(test)]
mod tests {
    use super::*;

    fn example_data() -> Vec<Range> {
        Day02::parse(
            "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124",
        )
    }

    #[test]
    fn part1_test() {
        assert_eq!(Day02::part1(&example_data()), 1227775554);
    }

    #[test]
    fn part2_test() {
        assert_eq!(Day02::part2(&example_data()), 4174379265);
    }

    #[test]
    fn part2_test_edge_case() {
        assert_eq!(Day02::part2(&vec![Range::new("1", "21")]), 11);
    }
}
