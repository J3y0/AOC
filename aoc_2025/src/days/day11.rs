use std::collections::HashMap;

use memoize::memoize;

use crate::days::Solution;

pub struct Day11;

pub type Rack = HashMap<String, Vec<String>>;

impl Solution for Day11 {
    type Input = Rack;

    fn parse(data: &str) -> Self::Input {
        let mut map = HashMap::new();
        for l in data.lines() {
            let (inp, outs) = l.split_once(":").unwrap();
            let outs = outs.split_ascii_whitespace().map(String::from).collect();

            map.insert(String::from(inp), outs);
        }

        map
    }

    fn part1(input: &Self::Input) -> usize {
        dfs(input, "you".to_string(), "out".to_string())
    }

    fn part2(input: &Self::Input) -> usize {
        dfs_2(input, false, false, "svr".to_string(), "out".to_string())
    }
}

#[memoize(Ignore: rack)]
fn dfs(rack: &Rack, cur: String, end: String) -> usize {
    if cur == end {
        return 1;
    }

    match rack.get(&cur) {
        Some(children) => {
            let mut count = 0;
            for c in children {
                count += dfs(rack, c.clone(), end.clone());
            }

            count
        }
        None => 0,
    }
}

#[memoize(Ignore: rack)]
fn dfs_2(rack: &Rack, seen_fft: bool, seen_dac: bool, cur: String, end: String) -> usize {
    if seen_fft && seen_dac && cur == end {
        return 1;
    }

    match rack.get(&cur) {
        Some(children) => {
            let mut count = 0;
            for c in children {
                count += dfs_2(
                    rack,
                    seen_fft | (c == "fft"),
                    seen_dac | (c == "dac"),
                    c.clone(),
                    end.clone(),
                );
            }

            count
        }
        None => 0,
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn part1_test() {
        let input = Day11::parse(
            "aaa: you hhh
you: bbb ccc
bbb: ddd eee
ccc: ddd eee fff
ddd: ggg
eee: out
fff: out
ggg: out
hhh: ccc fff iii
iii: out",
        );
        assert_eq!(Day11::part1(&input), 5);
    }

    #[test]
    fn part2_test() {
        let input = Day11::parse(
            "svr: aaa bbb
aaa: fft
fft: ccc
bbb: tty
tty: ccc
ccc: ddd eee
ddd: hub
hub: fff
eee: dac
dac: fff
fff: ggg hhh
ggg: out
hhh: out",
        );
        assert_eq!(Day11::part2(&input), 2);
    }
}
