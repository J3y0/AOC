use diesel::prelude::*;
use diesel::{ConnectionResult, SqliteConnection};
use std::path::Path;
use std::{env, fmt, fs, io};

use super::model::Cookies;

pub const CONFIG_DIR: &str = ".config/aoc-helper";
pub const SESSION_FILE: &str = "session";

#[derive(Debug)]
pub enum SessionError {
    ConnectionError(diesel::result::ConnectionError),
    QueryError(diesel::result::Error),
    NotFoundError,
}

impl From<diesel::result::ConnectionError> for SessionError {
    fn from(value: diesel::result::ConnectionError) -> Self {
        Self::ConnectionError(value)
    }
}

impl From<diesel::result::Error> for SessionError {
    fn from(value: diesel::result::Error) -> Self {
        Self::QueryError(value)
    }
}

impl fmt::Display for SessionError {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> Result<(), fmt::Error> {
        match self {
            Self::ConnectionError(e) => write!(f, "connection to database error: {e}"),
            Self::QueryError(e) => write!(f, "query database error: {e}"),
            Self::NotFoundError => write!(f, "'.adventofcode.com' session not found"),
        }
    }
}

fn connect_database<P: AsRef<Path>>(path: P) -> ConnectionResult<SqliteConnection> {
    let database_url = path.as_ref().to_str().unwrap();
    SqliteConnection::establish(database_url)
}

pub fn retrieve_session<P: AsRef<Path>>(path: P) -> Result<String, SessionError> {
    use super::schema::moz_cookies::dsl::{host, moz_cookies, name};

    let mut conn = connect_database(path)?;
    println!("[+] connected to database");

    let records = moz_cookies
        .filter(host.eq(".adventofcode.com"))
        .filter(name.eq("session"))
        .limit(1)
        .load::<Cookies>(&mut conn)?;

    if !records.is_empty() {
        let session = &records[0];
        session
            .value
            .to_owned()
            .ok_or_else(|| SessionError::NotFoundError)
    } else {
        Err(SessionError::NotFoundError)
    }
}

pub fn write_session_to_file(session: &str, filename: &str) -> io::Result<()> {
    let home_path = env::home_dir().unwrap();
    let config_dir_path = home_path.join(CONFIG_DIR);
    if !fs::exists(&config_dir_path)? {
        fs::create_dir(&config_dir_path)?;
    }
    fs::write(config_dir_path.join(filename), session)?;
    Ok(())
}
