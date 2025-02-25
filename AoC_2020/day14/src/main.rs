use regex::Regex;
use std::{fs, time::Instant};

#[derive(Debug)]
pub struct Mask {
    and_mask: u64,
    or_mask: u64,
}

#[derive(Debug)]
pub enum Instruction {
    UpdateMask(Mask),
    UpdateMemory((usize, u64)),
}

pub struct Program {
    instr: Vec<Instruction>,
    size_mem: usize,
}

fn main() {
    let content =
        fs::read_to_string("./input/day14.txt").expect("Something went wrong reading the file");

    let program = parse(&content);

    let time = Instant::now();
    let p1 = part1(&program);
    let p2 = part2(&program);
    let duration = time.elapsed();
    println!("Part1: {}", p1);
    println!("Part2: {}", p2);

    println!("Time: {} ns", duration.as_nanos());
}

pub fn part1(program: &Program) -> u64 {
    let mut memory = vec![0; program.size_mem + 1];
    let mut current_mask = Mask {
        and_mask: 0,
        or_mask: 0,
    };

    for instr in &program.instr {
        match instr {
            Instruction::UpdateMask(new_mask) => {
                current_mask.and_mask = new_mask.and_mask;
                current_mask.or_mask = new_mask.or_mask;
            }
            Instruction::UpdateMemory((loc, val)) => {
                let mask_val = (val | current_mask.or_mask) & current_mask.and_mask;
                memory[*loc] = mask_val;
            }
        }
    }

    memory.iter().sum()
}

#[allow(unused)]
pub fn part2(program: &Program) -> usize {
    0
}

fn parse(input: &str) -> Program {
    let mut size_mem = 0;
    let mut instr = vec![];
    for line in input.lines() {
        if line.starts_with("mask = ") {
            let mut and_mask: u64 = 0xffffffffff;
            let mut or_mask: u64 = 0x0;
            line[7..].chars().enumerate().for_each(|(i, c)| match c {
                '1' => or_mask |= 1 << (35 - i),
                '0' => and_mask ^= 1 << (35 - i),
                'X' => (),
                _ => unreachable!(),
            });

            instr.push(Instruction::UpdateMask(Mask { and_mask, or_mask }));
        } else {
            let re = Regex::new(r"mem\[(\d+)\] = (\d+)").unwrap();

            if let Some(cap) = re.captures(line) {
                let (_, [loc, val]) = cap.extract();
                let parsed_loc = loc.parse().unwrap();

                size_mem = size_mem.max(parsed_loc);

                instr.push(Instruction::UpdateMemory((
                    parsed_loc,
                    val.parse().unwrap(),
                )));
            }
        }
    }

    Program { instr, size_mem }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        let p1 = parse(
            "mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0",
        );
        assert_eq!(part1(&p1), 165);
    }

    #[test]
    fn test_part2() {
        let p2 = parse(
            "mask = 000000000000000000000000000000X1001X
mem[42] = 100
mask = 00000000000000000000000000000000X0XX
mem[26] = 1",
        );
        assert_eq!(part2(&p2), 208);
    }
}
