use std::{fs, mem, time::Instant};
use utils::{grid::Grid, point::Point};

struct Seat {
    index: usize,
    neighbors: Vec<usize>,
}

const FLOOR: u8 = b'.';
const EMPTY: u8 = b'L';
const OCCUPIED: u8 = b'#';

const DIRECTIONS: [Point; 8] = [
    Point::new(-1, 0),
    Point::new(1, 0),
    Point::new(0, -1),
    Point::new(0, 1),
    Point::new(-1, 1),
    Point::new(1, -1),
    Point::new(-1, -1),
    Point::new(1, 1),
];

fn main() {
    let content =
        fs::read_to_string("./input/day11.txt").expect("Something went wrong reading the file");

    let seats = parse(&content);

    let time = Instant::now();
    let p1 = part1(&seats);
    let p2 = part2(&seats);
    let duration = time.elapsed();
    println!("Part1: {}", p1);
    println!("Part2: {}", p2);

    println!("Time: {} ns", duration.as_nanos());
}

pub fn part1(seats: &Grid) -> usize {
    simulate_people(seats, true, 4)
}

pub fn part2(seats: &Grid) -> usize {
    simulate_people(seats, false, 5)
}

fn parse(input: &str) -> Grid {
    Grid::parse(input)
}

fn simulate_people(grid: &Grid, part_one: bool, value: usize) -> usize {
    let mut seats: Vec<Seat> = Vec::new();
    for i in 0..grid.height {
        for j in 0..grid.width {
            let pos = Point::new(i as i32, j as i32);
            if grid[pos] == FLOOR {
                continue;
            }

            let mut seat = Seat {
                index: (i * grid.width + j) as usize,
                neighbors: Vec::new(),
            };

            for d in DIRECTIONS {
                let next = pos + d;
                if part_one {
                    if !grid.outside(&next) && grid[next] != FLOOR {
                        seat.neighbors.push((next.x * grid.width + next.y) as usize);
                    }
                } else {
                    let mut next = pos + d;
                    while !grid.outside(&next) {
                        if grid[next] != FLOOR {
                            seat.neighbors.push((next.x * grid.width + next.y) as usize);
                            break;
                        }
                        next = next + d;
                    }
                }
            }
            seats.push(seat);
        }
    }

    let mut current = vec![EMPTY; grid.len()];
    let mut next = vec![0; grid.len()];
    let mut change = true;

    while change {
        change = false;

        for seat in &seats {
            // Count neighbors
            let mut nb_n = 0;
            for ineigh in &seat.neighbors {
                if current[*ineigh] == OCCUPIED {
                    nb_n += 1;
                }
            }

            if current[seat.index] == EMPTY && nb_n == 0 {
                next[seat.index] = OCCUPIED;
                change |= true;
            } else if current[seat.index] == OCCUPIED && nb_n >= value {
                next[seat.index] = EMPTY;
                change |= true;
            } else {
                next[seat.index] = current[seat.index];
            }
        }
        mem::swap(&mut current, &mut next);
    }

    current.iter().filter(|&s| b'#'.eq(s)).count()
}

#[cfg(test)]
mod tests {
    use super::*;

    fn example_data() -> Grid {
        parse(
            "L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL",
        )
    }

    #[test]
    fn test_part1() {
        assert_eq!(part1(&example_data()), 37);
    }

    #[test]
    fn test_part2() {
        assert_eq!(part2(&example_data()), 26);
    }
}
