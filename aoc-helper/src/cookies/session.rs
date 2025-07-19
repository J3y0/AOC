use diesel::{ConnectionResult, SqliteConnection};
use diesel::{prelude::*, result};
use std::fmt;
use std::path::Path;

use super::model::Cookies;

#[derive(Debug)]
pub enum SessionError {
    ConnectionError(result::ConnectionError),
    QueryError(result::Error),
    NotFoundError,
}

impl fmt::Display for SessionError {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> Result<(), fmt::Error> {
        match self {
            Self::ConnectionError(e) => write!(f, "connection to database error: {e}"),
            Self::QueryError(e) => write!(f, "query error: {e}"),
            Self::NotFoundError => write!(f, ".adventofcode.com session not found"),
        }
    }
}

pub fn retrieve_session(path: impl AsRef<Path>) -> Result<String, SessionError> {
    use super::schema::moz_cookies::dsl::{host, moz_cookies, name};

    let mut conn = connect_database(path).map_err(|e| SessionError::ConnectionError(e))?;
    println!("[+] connected to database");

    let records = moz_cookies
        .filter(host.eq(".adventofcode.com"))
        .filter(name.eq("session"))
        .limit(1)
        .load::<Cookies>(&mut conn)
        .map_err(|e| SessionError::QueryError(e))?;

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

fn connect_database(path: impl AsRef<Path>) -> ConnectionResult<SqliteConnection> {
    let database_url = path.as_ref().to_str().unwrap();
    SqliteConnection::establish(database_url)
}
