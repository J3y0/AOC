mod client;
mod commands;
mod cookies;

use clap::{Args, Parser, Subcommand, ValueEnum};
use commands::{cmd_get_input_file, cmd_get_session, cmd_set_session, cmd_submit_answer};
use std::{fmt::Display, ops::RangeInclusive, process};

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

impl Display for Part {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        match self {
            Self::One => write!(f, "1"),
            Self::Two => write!(f, "2"),
        }
    }
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
    Input(InputArgs),
    /// Submit an answer
    Answer(AnswerArgs),
    /// Retrieve cookie session from Firefox
    GetSession,
    /// Set given cookie session
    SetSession { session: String },
}

#[derive(Args)]
struct InputArgs {
    #[arg(short, long, help = "output filepath [default: <year>_<day>_day.txt]")]
    output: Option<String>,
    #[arg(short, long, help = "year to use", value_parser = year_in_range)]
    year: usize,
    #[arg(short, long, help = "day to use", value_parser = day_in_range)]
    day: usize,
}

#[derive(Args)]
struct AnswerArgs {
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

    run(&cli).unwrap_or_else(|err| {
        eprintln!("error: {err:?}");
        process::exit(1);
    });
}

fn run(opts: &Cli) -> anyhow::Result<()> {
    match &opts.command {
        Command::Input(input_args) => {
            let output = match input_args.output.clone() {
                Some(o) => o,
                None => format!("{}_{}_day.txt", input_args.year, input_args.day),
            };

            cmd_get_input_file(input_args.year, input_args.day, &output)?;
        }
        Command::Answer(answer_args) => {
            cmd_submit_answer(
                answer_args.year,
                answer_args.day,
                &answer_args.part,
                &answer_args.answer,
            )?;
        }
        Command::GetSession => cmd_get_session()?,
        Command::SetSession { session } => cmd_set_session(session)?,
    }
    Ok(())
}
