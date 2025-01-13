use std::{fs, time::Instant};

#[derive(Debug)]
pub enum Instruction {
    Nop(i32),
    Acc(i32),
    Jmp(i32),
}

fn main() {
    let content =
        fs::read_to_string("./input/day08.txt").expect("Something went wrong reading the file");

    let program = parse(&content);

    let time = Instant::now();
    let p1 = part1(&program);
    let p2 = part2(&program);
    let duration = time.elapsed();
    println!("Part1: {}", p1);
    println!("Part2: {}", p2);

    println!("Time: {} ns", duration.as_nanos());
}

pub fn part1(input: &[Instruction]) -> i32 {
    let seen = &mut vec![false; input.len()];
    run_program(input, 0, 0, seen, 1).unwrap()
}

pub fn part2(input: &[Instruction]) -> i32 {
    let mut acc = 0;
    let mut ip = 0;
    let seen = &mut vec![false; input.len()];
    while ip < input.len() {
        match input[ip] {
            Instruction::Acc(val) => {
                ip += 1;
                acc += val;
            }
            Instruction::Jmp(off) => {
                // Do a Nop
                if let Some(res) = run_program(input, acc, ip + 1, seen, 2) {
                    acc = res;
                    break;
                }
                ip = (ip as i32 + off) as usize;
            }
            Instruction::Nop(off) => {
                // Do a jmp
                if let Some(res) = run_program(input, acc, (ip as i32 + off) as usize, seen, 2) {
                    acc = res;
                    break;
                }
                ip += 1;
            }
        }
    }

    acc
}

fn run_program(
    program: &[Instruction],
    mut acc: i32,
    mut ip: usize,
    seen: &mut [bool],
    part: u8,
) -> Option<i32> {
    while ip < program.len() {
        if seen[ip] {
            if part == 1 {
                break;
            } else {
                return None;
            }
        }

        seen[ip] = true;

        match program[ip] {
            Instruction::Acc(val) => {
                ip += 1;
                acc += val;
            }
            Instruction::Jmp(offset) => ip = (ip as i32 + offset) as usize,
            Instruction::Nop(_) => ip += 1,
        }
    }

    Some(acc)
}

pub fn parse(input: &str) -> Vec<Instruction> {
    let mut instructions = Vec::new();
    for l in input.lines() {
        let mut iter = l.split_ascii_whitespace();

        let op = iter.next().unwrap();
        let operand = iter.next().unwrap().parse::<i32>().unwrap();

        match op {
            "acc" => instructions.push(Instruction::Acc(operand)),
            "nop" => instructions.push(Instruction::Nop(operand)),
            "jmp" => instructions.push(Instruction::Jmp(operand)),
            _ => unreachable!(),
        }
    }

    instructions
}

#[cfg(test)]
mod tests {
    use super::*;

    fn example_data() -> Vec<Instruction> {
        parse(
            "nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6",
        )
    }

    #[test]
    fn test_part1() {
        assert_eq!(part1(&example_data()), 5);
    }

    #[test]
    fn test_part2() {
        assert_eq!(part2(&example_data()), 8);
    }
}
