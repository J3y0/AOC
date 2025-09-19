mod cli;
mod client;
mod commands;
mod cookies;
mod logging;

use clap::Parser;
use cli::Command;
use commands::{cmd_get_input_file, cmd_get_session, cmd_set_session, cmd_submit_answer};
use log::{LevelFilter, error};
use std::process;

const LOG_FILE: &str = "aoc.log";

fn main() {
    let cli = cli::Cli::parse();

    let log_level = if cli.verbose {
        LevelFilter::Debug
    } else {
        LevelFilter::Info
    };

    if let Err(e) = logging::init_logs(log_level, LOG_FILE) {
        error!("failed to init logs: {e}");
        process::exit(1);
    }

    run(&cli).unwrap_or_else(|err| {
        error!("{err:?}");
        process::exit(1);
    });
}

fn run(opts: &cli::Cli) -> anyhow::Result<()> {
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
