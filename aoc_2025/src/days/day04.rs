use crate::days::Solution;

pub struct Day04;

pub type Grid = Vec<Vec<char>>;

#[derive(Debug)]
pub struct Map {
    grid: Grid,
    rolls: Vec<(usize, usize)>,
}

impl Solution for Day04 {
    type Input = Map;

    fn parse(data: &str) -> Self::Input {
        let mut rolls = vec![];
        let mut grid = vec![];
        for (i, l) in data.lines().enumerate() {
            let mut chars = Vec::with_capacity(l.len());
            for (j, c) in l.chars().enumerate() {
                chars.push(c);
                if c == '@' {
                    rolls.push((i, j));
                }
            }

            grid.push(chars);
        }

        Map { grid, rolls }
    }

    fn part1(input: &Self::Input) -> usize {
        input
            .rolls
            .iter()
            .filter(|r| is_accessible(r, &input.grid))
            .count()
    }

    fn part2(input: &Self::Input) -> usize {
        let mut still = true;
        let mut tot = 0;

        let mut grid = input.grid.clone();
        let mut cur_rolls = input.rolls.clone();

        while still {
            let mut next_rolls = vec![];
            still = false;

            for r in &cur_rolls {
                if is_accessible(r, &grid) {
                    still = true;
                    tot += 1;

                    // Remove the roll
                    grid[r.0][r.1] = '.';
                } else {
                    next_rolls.push(*r);
                }
            }

            cur_rolls.clear();
            cur_rolls.extend_from_slice(&next_rolls);
        }

        tot
    }
}

fn is_accessible(pos: &(usize, usize), grid: &Grid) -> bool {
    let neighbors = get_valid_neighbors(pos, grid);

    neighbors.iter().filter(|n| grid[n.0][n.1] == '@').count() < 4
}

fn get_valid_neighbors(pos: &(usize, usize), grid: &Grid) -> Vec<(usize, usize)> {
    let h = grid.len();
    let w = grid.first().unwrap_or(&vec![]).len();

    let mut n = vec![];
    for nx in pos.0.saturating_sub(1)..=(pos.0 + 1) {
        for ny in pos.1.saturating_sub(1)..=(pos.1 + 1) {
            if nx >= h || ny >= w {
                continue;
            }

            if (nx, ny) == *pos {
                continue;
            }

            n.push((nx, ny));
        }
    }

    n
}

#[cfg(test)]
mod tests {
    use super::*;

    fn example_data() -> Map {
        Day04::parse(
            "..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.",
        )
    }

    #[test]
    fn part1_test() {
        assert_eq!(Day04::part1(&example_data()), 13);
    }

    #[test]
    fn part2_test() {
        assert_eq!(Day04::part2(&example_data()), 43);
    }
}
