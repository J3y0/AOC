mod commands;
mod cookies;

use clap::{Args, Parser, Subcommand, ValueEnum};
use commands::{cmd_get_session, cmd_set_session};
use std::ops::RangeInclusive;

const YEAR_RANGE: RangeInclusive<usize> = 2015..=2025;
const DAY_RANGE: RangeInclusive<usize> = 1..=25;

fn year_in_range(y: &str) -> Result<usize, String> {
    let year: usize = y.parse().map_err(|_| format!("`{y}` is not a number."))?;
    if YEAR_RANGE.contains(&year) {
        Ok(year)
    } else {
        Err(format!(
            "`{year}` not in range {}-{}.",
            YEAR_RANGE.start(),
            YEAR_RANGE.end()
        ))
    }
}

fn day_in_range(d: &str) -> Result<usize, String> {
    let day: usize = d.parse().map_err(|_| format!("`{d}` is not a number."))?;
    if DAY_RANGE.contains(&day) {
        Ok(day)
    } else {
        Err(format!(
            "`{day}` not in range {}-{}.",
            DAY_RANGE.start(),
            DAY_RANGE.end()
        ))
    }
}

#[derive(Clone, ValueEnum)]
enum Part {
    One,
    Two,
}

#[derive(Parser)]
#[command(name = "aoc-helper")]
#[command(about = "CLI to help you interact with Advent Of Code.")]
struct Cli {
    #[command(subcommand)]
    command: Command,
}

#[derive(Subcommand)]
enum Command {
    /// Get an input file
    Get(GetArgs),
    /// Submit an answer
    Submit(SubmitArgs),
    /// Retrieve cookie session and set it. You should have previously logged in adventofcode using Firefox.
    GetSession,
    /// Set specified cookie session.
    SetSession { session: String },
}

#[derive(Args)]
struct GetArgs {
    #[arg(short, long, help = "output filepath [default: <year>_<day>_day.txt]")]
    output: Option<String>,
    #[arg(short, long, help = "year to use", value_parser = year_in_range)]
    year: usize,
    #[arg(short, long, help = "day to use", value_parser = day_in_range)]
    day: usize,
}

#[derive(Args)]
struct SubmitArgs {
    #[arg(short, long, help = "year to use for submission", value_parser = year_in_range)]
    year: usize,
    #[arg(short, long, help = "day to use for submission", value_parser = day_in_range)]
    day: usize,
    #[arg(short, long, help = "part to use for submission", value_enum)]
    part: Part,
    answer: String,
}

fn main() {
    let cli = Cli::parse();

    match cli.command {
        Command::Get(get_args) => {
            let output = match get_args.output {
                Some(o) => o,
                None => format!("{}_{}_day.txt", get_args.year, get_args.day),
            };

            println!(
                "year: {}, day: {}, output: {}",
                get_args.year, get_args.day, output
            );
        }
        Command::Submit(submit_args) => {
            println!(
                "year: {}, day: {}, part: {}",
                submit_args.year, submit_args.day, submit_args.part
            );
        }
        Command::GetSession => cmd_get_session(),
        Command::SetSession { session } => cmd_set_session(&session),
    }
}
