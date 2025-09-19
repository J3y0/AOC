use clap::{Args, Parser, Subcommand, ValueEnum};
use std::fmt::Display;
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
pub enum Part {
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
pub struct Cli {
    #[arg(short, long, global = true, help = "debug output")]
    pub verbose: bool,
    #[command(subcommand)]
    pub command: Command,
}

#[derive(Subcommand)]
pub enum Command {
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
pub struct InputArgs {
    #[arg(short, long, help = "output filepath [default: <year>_<day>_day.txt]")]
    pub output: Option<String>,
    #[arg(short, long, help = "year to use", value_parser = year_in_range)]
    pub year: usize,
    #[arg(short, long, help = "day to use", value_parser = day_in_range)]
    pub day: usize,
}

#[derive(Args)]
pub struct AnswerArgs {
    #[arg(short, long, help = "year to use for submission", value_parser = year_in_range)]
    pub year: usize,
    #[arg(short, long, help = "day to use for submission", value_parser = day_in_range)]
    pub day: usize,
    #[arg(short, long, help = "part to use for submission", value_enum)]
    pub part: Part,
    pub answer: String,
}
