use crate::days::Solution;

pub struct Day03;

pub type Bank = Vec<u8>;

impl Solution for Day03 {
    type Input = Vec<Bank>;

    fn parse(data: &str) -> Self::Input {
        data.lines()
            .map(|line| {
                line.chars()
                    .map(|c| c.to_digit(10).unwrap() as u8)
                    .collect()
            })
            .collect()
    }

    fn part1(input: &Self::Input) -> usize {
        let mut tot = 0;
        for bank in input {
            tot += max_jolt(&bank, 2);
        }

        tot
    }

    fn part2(input: &Self::Input) -> usize {
        let mut tot = 0;
        for bank in input {
            tot += max_jolt(&bank, 12);
        }

        tot
    }
}

fn max_jolt(bank: &Bank, nb_bat: usize) -> usize {
    let mut res = 0;
    let mut pos = 0;
    let n = bank.len();

    for i in (1..=nb_bat).rev() {
        // enough numbers should remain to compose the result
        let end = n - i;
        let mut max = 0;
        for j in pos..=end {
            if bank[j] > max {
                max = bank[j];
                pos = j + 1;
            }
        }

        res += (max as usize) * 10_usize.pow((i - 1) as u32);
    }

    res
}

#[cfg(test)]
mod tests {
    use super::*;

    fn example_data() -> Vec<Bank> {
        Day03::parse(
            "987654321111111
811111111111119
234234234234278
818181911112111",
        )
    }

    #[test]
    fn part1_test() {
        assert_eq!(Day03::part1(&example_data()), 357);
    }

    #[test]
    fn part2_test() {
        assert_eq!(Day03::part2(&example_data()), 3121910778619);
    }
}
