use AoC_2025::{
    ansi::{BOLD, RED, RESET, YELLOW},
    cli,
    days::{self, Solution},
};
use clap::Parser;
use std::{fs::read_to_string, process, time::Instant};

fn run<S>(day: usize)
where
    S: Solution,
{
    let path = format!("input/{day:02}_day.txt");
    if let Ok(data) = read_to_string(&path) {
        let input = S::parse(&data);

        let now = Instant::now();
        let part1 = S::part1(&input);
        let elapsed_part1 = now.elapsed();

        let now = Instant::now();
        let part2 = S::part2(&input);
        let elapsed_part2 = now.elapsed();

        println!("{YELLOW}------------{RESET}");
        println!("{BOLD}{YELLOW}   Day {day:02}{RESET}");
        println!("{YELLOW}------------{RESET}");
        println!("Part1: {part1}");
        println!("Time1: {} ns", elapsed_part1.as_nanos());
        println!("Part2: {part2}");
        println!("Time2: {} ns", elapsed_part2.as_nanos());
    } else {
        println!("{RED}------------{RESET}");
        println!("{BOLD}{RED}   Day {day:02}{RESET}");
        println!("{RED}------------{RESET}");
        println!("Cannot read input at \"{path}\"");
    }
}

fn main() {
    let cli = cli::Cli::parse();

    match cli.day {
        1 => run::<days::Day01>(cli.day),
        2 => run::<days::Day02>(cli.day),
        3 => run::<days::Day03>(cli.day),
        4 => run::<days::Day04>(cli.day),
        5 => run::<days::Day05>(cli.day),
        6 => run::<days::Day06>(cli.day),
        7 => run::<days::Day07>(cli.day),
        8 => run::<days::Day08>(cli.day),
        9 => run::<days::Day09>(cli.day),
        10 => run::<days::Day10>(cli.day),
        11 => run::<days::Day11>(cli.day),
        _ => {
            eprintln!("Day {:02} is not implemented yet!", cli.day);
            process::exit(1);
        }
    };
}
