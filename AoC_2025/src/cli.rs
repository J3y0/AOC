use clap::Parser;
use std::ops::RangeInclusive;

const DAY_RANGE: RangeInclusive<usize> = 1..=25;
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

#[derive(Parser)]
pub struct Cli {
    #[arg(help = "day to run", value_parser = day_in_range)]
    pub day: usize,
}
