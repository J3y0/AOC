use AoC_2025::{
    ansi::{BOLD, RED, RESET, YELLOW},
    cli,
    days::{self, Solution},
};
use clap::Parser;
use std::{fs::read_to_string, process, time::Instant};

fn run(day: usize, sol: impl Solution) {
    let path = format!("input/{day:02}_day.txt");
    if let Ok(data) = read_to_string(&path) {
        let now = Instant::now();
        let input = sol.parse(&data);
        let part1 = sol.part1(&input);
        let part2 = sol.part2(&input);
        let elapsed = now.elapsed();

        println!("{YELLOW}------------{RESET}");
        println!("{BOLD}{YELLOW}   Day {day:02}{RESET}");
        println!("{YELLOW}------------{RESET}");
        println!("Part1: {part1}");
        println!("Part2: {part2}");
        println!("Time: {} ns", elapsed.as_nanos());
    } else {
        println!("{RED}------------{RESET}");
        println!("{BOLD}{RED}   Day {day:02}{RESET}");
        println!("{RED}------------{RESET}");
        println!("Cannot read input at \"{path}\"");
    }
}

fn main() {
    let cli = cli::Cli::parse();

    let sol = match cli.day {
        1 => days::day01::Day01,
        _ => {
            eprintln!("Day {:02} is not implemented yet!", cli.day);
            process::exit(1);
        }
    };

    run(cli.day, sol);
}
