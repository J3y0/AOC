use std::collections::{HashMap, HashSet, VecDeque};

use crate::days::Solution;

pub struct Day07;

pub struct Tachyon {
    start: (usize, usize),
    height: usize,
    grid: Vec<Vec<u8>>,
}

impl Solution for Day07 {
    type Input = Tachyon;

    fn parse(data: &str) -> Self::Input {
        let grid: Vec<Vec<u8>> = data.lines().map(|l| l.as_bytes().to_vec()).collect();
        let height = grid.len();
        let start = (0, grid[0].iter().position(|&c| c == b'S').unwrap());

        Tachyon {
            height,
            grid,
            start,
        }
    }

    fn part1(input: &Self::Input) -> usize {
        let mut splitted = 0;
        let mut seen: HashSet<(usize, usize)> = HashSet::new();
        let mut beams: VecDeque<(usize, usize)> = VecDeque::new();

        beams.push_back(input.start);
        seen.insert(input.start);

        while let Some(cur) = beams.pop_front() {
            let newpos = (cur.0 + 1, cur.1);

            // beam reached end
            if newpos.0 >= input.height {
                continue;
            }

            let new_tile = input.grid[newpos.0][newpos.1];
            match new_tile {
                b'^' => {
                    let left = (newpos.0, newpos.1 - 1);
                    if seen.insert(left) {
                        beams.push_back(left);
                    }
                    let right = (newpos.0, newpos.1 + 1);
                    if seen.insert(right) {
                        beams.push_back(right);
                    }
                    splitted += 1;
                }
                b'.' => {
                    if seen.insert(newpos) {
                        beams.push_back(newpos);
                    }
                }
                _ => {}
            }
        }

        splitted
    }

    fn part2(input: &Self::Input) -> usize {
        let mut timelines = 0;
        let mut beams: VecDeque<(usize, usize)> = VecDeque::new();
        let mut count: HashMap<(usize, usize), usize> = HashMap::new();
        beams.push_back(input.start);
        count.insert(input.start, 1);

        while let Some(pos) = beams.pop_front() {
            let newpos = (pos.0 + 1, pos.1);
            let cur_count = count[&pos];

            if newpos.0 >= input.height {
                timelines += cur_count;
                continue;
            }

            let new_tile = input.grid[newpos.0][newpos.1];
            match new_tile {
                b'^' => {
                    let l = (newpos.0, newpos.1 - 1);
                    let entry = count.entry(l).or_insert(0);
                    if *entry == 0 {
                        beams.push_back(l);
                    }
                    *entry += cur_count;

                    let r = (newpos.0, newpos.1 + 1);
                    let entry = count.entry(r).or_insert(0);
                    if *entry == 0 {
                        beams.push_back(r);
                    }
                    *entry += cur_count;
                }
                b'.' => {
                    let entry = count.entry(newpos).or_insert(0);
                    if *entry == 0 {
                        beams.push_back(newpos);
                    }
                    *entry += cur_count;
                }
                _ => {}
            }
        }

        timelines
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    fn example_data() -> Tachyon {
        Day07::parse(
            ".......S.......
...............
.......^.......
...............
......^.^......
...............
.....^.^.^.....
...............
....^.^...^....
...............
...^.^...^.^...
...............
..^...^.....^..
...............
.^.^.^.^.^...^.
...............",
        )
    }

    #[test]
    fn part1_test() {
        assert_eq!(Day07::part1(&example_data()), 21);
    }

    #[test]
    fn part2_test() {
        assert_eq!(Day07::part2(&example_data()), 40);
    }
}
