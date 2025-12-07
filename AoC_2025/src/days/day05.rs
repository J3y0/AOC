use crate::days::Solution;

pub struct Day05;

#[derive(Clone, Debug)]
struct Range {
    start: u64,
    end: u64,
}

impl Range {
    fn len(&self) -> u64 {
        self.end - self.start + 1
    }

    fn contains(&self, val: &u64) -> bool {
        (self.start..=self.end).contains(val)
    }

    fn overlap(&self, r: &Range) -> bool {
        self.contains(&r.start)
            || self.contains(&r.end)
            || r.contains(&self.start)
            || r.contains(&self.end)
    }

    /// merge two ranges together, if they can be merged
    /// Return:
    ///     - None if ranges do not overlap
    ///     - The new merged range if they overlap
    fn merge(&self, r: &Range) -> Option<Range> {
        if !self.overlap(r) {
            return None;
        }

        let new_s = self.start.min(r.start);
        let new_e = self.end.max(r.end);

        Some(Range {
            start: new_s,
            end: new_e,
        })
    }
}

impl From<&str> for Range {
    fn from(value: &str) -> Self {
        let (start, end) = value.split_once('-').unwrap();
        let start_parsed = start.parse().unwrap();
        let end_parsed = end.parse().unwrap();

        Range {
            start: start_parsed,
            end: end_parsed,
        }
    }
}

#[derive(Debug)]
pub struct Kitchen {
    fresh: Vec<Range>,
    available: Vec<u64>,
}

impl Solution for Day05 {
    type Input = Kitchen;

    fn parse(data: &str) -> Self::Input {
        let (ranges, available) = data.split_once("\n\n").unwrap();

        let fresh = ranges.lines().map(|l| Range::from(l)).collect();

        let available = available
            .lines()
            .map(|l| l.parse::<u64>().unwrap())
            .collect();

        Kitchen { fresh, available }
    }

    fn part1(input: &Self::Input) -> usize {
        input
            .available
            .iter()
            .filter(|&i| is_fresh(i, &input.fresh))
            .count()
    }

    fn part2(input: &Self::Input) -> usize {
        let mut rgs = input.fresh.clone();

        // O(n*log(n))
        rgs.sort_by(|a, b| a.start.cmp(&b.start));

        let mut merged_rgs = vec![rgs[0].clone()];
        for i in 1..rgs.len() {
            let rg = &rgs[i];

            let last = merged_rgs.len() - 1;
            match merged_rgs[last].merge(rg) {
                Some(merged) => merged_rgs[last] = merged,
                None => merged_rgs.push(rg.clone()),
            }
        }

        merged_rgs.iter().map(|r| r.len() as usize).sum()
    }
}

fn is_fresh(ingredient: &u64, ranges: &[Range]) -> bool {
    ranges.iter().any(|r| r.contains(ingredient))
}

#[cfg(test)]
mod tests {
    use super::*;

    fn example_data() -> Kitchen {
        Day05::parse(
            "3-5
10-14
16-20
12-18

1
5
8
11
17
32",
        )
    }

    #[test]
    fn part1_test() {
        assert_eq!(Day05::part1(&example_data()), 3);
    }

    #[test]
    fn part2_test() {
        assert_eq!(Day05::part2(&example_data()), 14);
    }
}
