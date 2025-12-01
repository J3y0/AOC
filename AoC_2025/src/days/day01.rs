use crate::days::Solution;

pub struct Day01;

impl Solution for Day01 {
    type Input = Vec<i16>;

    fn parse(&self, data: &str) -> Self::Input {
        data.lines()
            .map(|line| {
                let dir = line.chars().nth(0).unwrap();
                let parsed = (&line[1..]).parse::<i16>().unwrap();
                if dir == 'L' { -parsed } else { parsed }
            })
            .collect()
    }

    fn part1(&self, input: &Self::Input) -> usize {
        // Start at position 50
        let mut cur = 50;
        let mut passwd = 0;
        for elt in input {
            cur = (cur + elt).rem_euclid(100);
            if cur == 0 {
                passwd += 1;
            }
        }

        passwd
    }

    fn part2(&self, input: &Self::Input) -> usize {
        // Start at position 50
        let mut cur = 50;
        let mut passwd = 0;
        for &elt in input {
            // add all additional complete rotations
            passwd += ((cur + elt) / 100).abs();

            // Add if point 0 by subtracting the value
            if elt < 0 && cur != 0 && -elt >= cur {
                passwd += 1;
            }

            cur = (cur + elt).rem_euclid(100);
        }

        passwd as usize
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    fn example_data() -> Vec<i16> {
        let sol = Day01 {};
        sol.parse(
            r#"L68
L30
R48
L5
R60
L55
L1
L99
R14
L82"#,
        )
    }

    #[test]
    fn part1_test() {
        let sol = Day01 {};
        assert_eq!(sol.part1(&example_data()), 3);
    }

    #[test]
    fn part2_test() {
        let sol = Day01 {};
        assert_eq!(sol.part2(&example_data()), 6);
    }
}
