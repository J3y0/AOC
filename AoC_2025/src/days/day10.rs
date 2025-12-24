use std::{collections::HashMap, u16};

use crate::days::Solution;

pub struct Day10;

#[derive(Debug)]
pub struct Machine {
    lights: u32,
    buttons: Vec<Vec<u8>>,
    joltages: Vec<u16>,
}

impl Solution for Day10 {
    type Input = Vec<Machine>;

    fn parse(data: &str) -> Self::Input {
        data.lines()
            .map(|line| {
                let mut lights = 0;
                let mut buttons = vec![];
                let mut joltages = vec![];
                for elt in line.split_ascii_whitespace() {
                    let inner = &elt[1..elt.len() - 1];
                    match elt.chars().next().unwrap() {
                        '[' => {
                            lights = inner.chars().enumerate().fold(0, |acc, (i, c)| {
                                let add = if c == '#' { 1 << i } else { 0 };
                                acc + add
                            });
                        }
                        '(' => {
                            buttons.push(inner.split(',').map(|i| i.parse().unwrap()).collect());
                        }
                        '{' => {
                            joltages = inner.split(',').map(|i| i.parse().unwrap()).collect();
                        }
                        _ => unreachable!(),
                    }
                }

                Machine {
                    buttons,
                    lights,
                    joltages,
                }
            })
            .collect()
    }
    fn part1(input: &Self::Input) -> usize {
        input.iter().map(min_presses).sum()
    }

    fn part2(input: &Self::Input) -> usize {
        let mut tot = 0;
        for f in input {
            let mut cache: HashMap<Vec<u16>, Option<usize>> = HashMap::new();
            tot += recurse(&f.buttons, &f.joltages, &mut cache).unwrap();
        }

        tot
    }
}

fn min_presses(machine: &Machine) -> usize {
    for i in 1..machine.buttons.len() {
        let combi = Combination::new(machine.buttons.clone(), i);
        for c in combi {
            let bin_buttons = get_binary_buttons(&c);

            let res = bin_buttons.iter().fold(0, |acc, b| acc ^ b);
            if res == machine.lights {
                // Return directly as it is sure to be the lowest solution
                return c.len();
            }
        }
    }

    // there is always a solution in given input
    unreachable!()
}

fn get_binary_buttons(buttons: &[Vec<u8>]) -> Vec<u32> {
    buttons
        .iter()
        .map(|b| b.iter().fold(0, |acc, e| acc + (1 << e)))
        .collect()
}

// https://www.reddit.com/r/adventofcode/comments/1pk87hl/2025_day_10_part_2_bifurcate_your_way_to_victory/
fn recurse(
    buttons: &Vec<Vec<u8>>,
    joltages: &Vec<u16>,
    cache: &mut HashMap<Vec<u16>, Option<usize>>,
) -> Option<usize> {
    if let Some(cached) = cache.get(joltages) {
        return *cached;
    }

    // no more lights [0 0 0 ...]
    if joltages.iter().all(|j| *j == 0) {
        cache.insert(joltages.clone(), Some(0));
        return Some(0);
    }

    let bin_joltages = get_binary_joltages(&joltages);
    let mut min_presses = None;
    // should start from 0 to handle cases where all joltages are already even
    for i in 0..=buttons.len() {
        let combi = Combination::new(buttons.clone(), i);
        for c in combi {
            let bin_buttons = get_binary_buttons(&c);
            let res_combi = bin_buttons.iter().fold(0, |acc, b| acc ^ b);
            if res_combi == bin_joltages {
                // go to next iteration if the combination exceeds any joltage value
                let new_joltages = match get_next_joltages(&joltages, &bin_buttons) {
                    Some(nj) => nj,
                    None => continue,
                };
                let rec_presses = recurse(buttons, &new_joltages, cache).map(|v| c.len() + 2 * v);
                min_presses = min_presses.min(rec_presses).or(min_presses).or(rec_presses);
            }
        }
    }

    cache.insert(joltages.clone(), min_presses);

    min_presses
}

fn get_next_joltages(joltages: &[u16], bin_buttons: &[u32]) -> Option<Vec<u16>> {
    let mut next_joltages = vec![0u16; joltages.len()];
    for i in 0..joltages.len() {
        let to_sub = bin_buttons.iter().filter(|&b| *b & (1 << i) != 0).count() as u16;
        // overflow should not occur
        match joltages[i].checked_sub(to_sub) {
            Some(nj) => next_joltages[i] = nj / 2,
            None => {
                return None;
            }
        }
    }

    Some(next_joltages)
}

fn get_binary_joltages(joltages: &[u16]) -> u32 {
    joltages
        .iter()
        .enumerate()
        .fold(0, |acc, (i, jolt)| acc + (((*jolt as u32) % 2) << i))
}

struct Combination<T> {
    elements: Vec<T>,
    r: usize,
    n: usize,
    indices: Vec<usize>,
    first: bool,
}

impl<T: Clone> Combination<T> {
    fn new<I>(elements: I, r: usize) -> Self
    where
        I: IntoIterator<Item = T>,
    {
        let elements: Vec<T> = elements.into_iter().collect();

        Self {
            indices: (0..r).collect(),
            n: elements.len(),
            first: true,
            elements,
            r,
        }
    }
}

// code back itertools.combinations() from python
impl<T: Clone> Iterator for Combination<T> {
    type Item = Vec<T>;
    fn next(&mut self) -> Option<Self::Item> {
        if self.r > self.n {
            return None;
        }

        if self.first {
            self.first = false;
            return Some(
                self.indices
                    .iter()
                    .map(|&i| self.elements[i].clone())
                    .collect(),
            );
        }

        let mut rightmost_i = None;
        for i in (0..self.r).rev() {
            if self.indices[i] != self.n + i - self.r {
                rightmost_i = Some(i);
                break;
            }
        }

        let rightmost_i = rightmost_i?;

        self.indices[rightmost_i] += 1;
        for j in rightmost_i + 1..self.r {
            self.indices[j] = self.indices[j - 1] + 1;
        }

        Some(
            self.indices
                .iter()
                .map(|&i| self.elements[i].clone())
                .collect(),
        )
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    fn example_data() -> Vec<Machine> {
        Day10::parse(
            "[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}
[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}
[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}",
        )
    }

    #[test]
    fn part1_test() {
        assert_eq!(Day10::part1(&example_data()), 7);
    }

    #[test]
    fn part2_test() {
        assert_eq!(Day10::part2(&example_data()), 33);
    }

    // #[test]
    // fn debug_recurse() {
    //     assert_eq!(
    //         recurse(&vec![vec![0, 1], vec![0, 2], vec![1, 2]], &vec![2, 2, 2]),
    //         Some(3)
    //     )
    // }

    // #[test]
    // fn debug_recurse_double() {
    //     assert_eq!(
    //         recurse(&vec![vec![0, 1], vec![0, 2], vec![1, 2]], &vec![4, 4, 4]),
    //         Some(6)
    //     )
    // }

    // #[test]
    // fn debug_recurse_double_add() {
    //     assert_eq!(
    //         recurse(&vec![vec![0, 1], vec![0, 2], vec![1, 2]], &vec![5, 5, 4]),
    //         Some(7)
    //     )
    // }
}
