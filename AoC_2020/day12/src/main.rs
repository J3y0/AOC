use std::{fs, time::Instant};
use utils::point::Point;

pub struct Instruction {
    code: char,
    dist: i32,
}

fn main() {
    let content =
        fs::read_to_string("./input/day12.txt").expect("Something went wrong reading the file");

    let instructions = parse(&content);

    let time = Instant::now();
    let p1 = part1(&instructions);
    let p2 = part2(&instructions);
    let duration = time.elapsed();
    println!("Part1: {}", p1);
    println!("Part2: {}", p2);

    println!("Time: {} ns", duration.as_nanos());
}

pub fn part1(instructions: &[Instruction]) -> usize {
    let mut cdir = Point::new(0, 1);
    let mut cpos = Point::new(0, 0);
    for inst in instructions {
        match inst.code {
            'N' => cpos.x -= inst.dist,
            'S' => cpos.x += inst.dist,
            'W' => cpos.y -= inst.dist,
            'E' => cpos.y += inst.dist,
            'R' => rotate_right(&mut cdir, inst.dist),
            'L' =>  rotate_right(&mut cdir, 360-inst.dist),
            'F' => {
                cpos.x += cdir.x * inst.dist;
                cpos.y += cdir.y * inst.dist;
            },
            _ => unreachable!(),
        }
    }

    cpos.manhattan_distance(&Point::new(0, 0)) as usize
}

pub fn part2(instructions: &[Instruction]) -> usize {
    let mut cwaypoint = Point::new(-1, 10);
    let mut cpos = Point::new(0, 0);
    for inst in instructions {
        match inst.code {
            'N' => cwaypoint.x -= inst.dist,
            'S' => cwaypoint.x += inst.dist,
            'W' => cwaypoint.y -= inst.dist,
            'E' => cwaypoint.y += inst.dist,
            'R' => rotate_right(&mut cwaypoint, inst.dist),
            'L' =>  rotate_right(&mut cwaypoint, 360-inst.dist),
            'F' => {
                cpos.x += cwaypoint.x * inst.dist;
                cpos.y += cwaypoint.y * inst.dist;
            },
            _ => unreachable!(),
        }
    }

    cpos.manhattan_distance(&Point::new(0, 0)) as usize
}

fn parse(input: &str) -> Vec<Instruction> {
    input
        .lines()
        .map(|l| {
            let of = l[1..].parse::<i32>().unwrap();
            Instruction{ code: l.chars().next().unwrap(), dist: of}
        })
        .collect()
}

fn rotate_right(direction: &mut Point, dist: i32) {
    match dist {
        90 => {
            let temp = direction.x;
            direction.x = direction.y;
            direction.y = -temp;
        },
        180 => {
            direction.x = -direction.x;
            direction.y = -direction.y;
        },
        270 => {
            let temp = direction.x;
            direction.x = -direction.y;
            direction.y = temp;
        },
        _ => unreachable!(),
    }
}


#[cfg(test)]
mod tests {
    use super::*;

    fn example_data() -> Vec<Instruction> {
        parse("F10
N3
F7
R90
F11")
    }

    #[test]
    fn test_part1() {
        assert_eq!(part1(&example_data()), 25);
    }

    #[test]
    fn test_part2() {
        assert_eq!(part2(&example_data()), 286);
    }
}