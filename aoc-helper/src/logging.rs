use std::{env::home_dir, fs::File};

use log::SetLoggerError;
use simplelog::{
    CombinedLogger, ConfigBuilder, SharedLogger, TermLogger, TerminalMode, WriteLogger,
};

use crate::cookies::session::CONFIG_DIR;

pub fn init_logs(log_level: log::LevelFilter, filename: &str) -> Result<(), SetLoggerError> {
    let mut filepath = home_dir().unwrap();
    filepath.push(CONFIG_DIR);
    filepath.push(filename);

    let logfile = File::create(filepath).unwrap();
    let config_writelogger = ConfigBuilder::new().set_time_format_rfc2822().build();
    let config_termlogger = ConfigBuilder::new()
        .set_time_level(log::LevelFilter::Off)
        .set_thread_level(log::LevelFilter::Off)
        .set_target_level(log::LevelFilter::Off)
        .set_max_level(log::LevelFilter::Debug)
        .build();

    let loggers: Vec<Box<dyn SharedLogger>> = vec![
        WriteLogger::new(log::LevelFilter::Debug, config_writelogger, logfile),
        TermLogger::new(
            log_level,
            config_termlogger,
            TerminalMode::Mixed,
            simplelog::ColorChoice::Auto,
        ),
    ];

    CombinedLogger::init(loggers)
}
