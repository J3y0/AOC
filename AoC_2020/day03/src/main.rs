use std::{fs, time::Instant};

fn main() {
    let content =
        fs::read_to_string("./input/day03.txt").expect("Something went wrong reading the file");
    let grid = parse(&content);

    let time = Instant::now();
    let p1 = part1(&grid);
    let p2 = part2(&grid);
    let duration = time.elapsed();
    println!("Part1: {}", p1);
    println!("Part2: {}", p2);

    println!("Time: {} ns", duration.as_nanos());
}

pub fn part1(grid: &[&[u8]]) -> usize {
    toboggan(grid, (1, 3))
}

pub fn part2(grid: &[&[u8]]) -> usize {
    toboggan(grid, (1, 1))
        * toboggan(grid, (1, 3))
        * toboggan(grid, (1, 5))
        * toboggan(grid, (1, 7))
        * toboggan(grid, (2, 1))
}

pub fn parse(input: &str) -> Vec<&[u8]> {
    let rows: Vec<_> = input.lines().map(|l| l.as_bytes()).collect();
    let mut res = Vec::with_capacity(rows.len());

    rows.iter().for_each(|&l| res.push(l));
    res
}

fn toboggan(grid: &[&[u8]], slope: (usize, usize)) -> usize {
    let mut pos = (0, 0);
    let h = grid.len();
    let w = grid[0].len();

    let mut trees = 0;
    while pos.0 < h {
        if grid[pos.0][pos.1] == b'#' {
            trees += 1;
        }
        pos.0 += slope.0;
        pos.1 = (pos.1 + slope.1) % w;
    }
    trees
}

#[cfg(test)]
mod tests {
    use super::*;

    fn example_data() -> Vec<&'static [u8]> {
        parse(
            "..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#",
        )
    }

    #[test]
    fn test_part1() {
        assert_eq!(part1(&example_data()), 7);
    }

    #[test]
    fn test_part2() {
        assert_eq!(part2(&example_data()), 336);
    }
}
