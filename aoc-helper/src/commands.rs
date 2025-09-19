use anyhow::{Context, anyhow};
use log::{debug, info};
use std::{
    env, fs,
    path::{Path, PathBuf},
};

use crate::cookies::session;
use crate::{Part, client::AocClient};

pub fn cmd_set_session(session: &str) -> anyhow::Result<()> {
    session::write_session_to_file(session, session::SESSION_FILE)
        .with_context(|| format!("could not set up session: {session}"))?;

    info!("session set up successfully to: {session}");
    Ok(())
}

pub fn cmd_get_session() -> anyhow::Result<()> {
    let home_path = env::home_dir().unwrap();
    let search_path = home_path.join(".mozilla/firefox/*.default-release/cookies.sqlite");
    let search_path_str = search_path.to_str().unwrap();

    let mut path = PathBuf::new();
    for entry in glob::glob(search_path_str)
        .with_context(|| format!("failed to read glob pattern: {search_path_str}"))?
        .flatten()
    {
        path = entry;
    }

    debug!("found glob pattern for {search_path_str}");

    let filename = path.file_name().ok_or(anyhow!(
        "could not find firefox database file: cookies.sqlite"
    ))?;

    // tmp database if firefox open
    let tmp_path = Path::new("/tmp").join(filename);
    debug!("tmp_path: {:?}", &tmp_path);
    fs::copy(path, &tmp_path)?;

    let session = session::retrieve_session(&tmp_path)
        .map_err(|err| anyhow!(err))
        .context("could not retrieve session")?;

    debug!("retrieve session ok");

    session::write_session_to_file(&session, session::SESSION_FILE)
        .context("could not save session")?;
    fs::remove_file(tmp_path)?;

    info!("session retrieved and saved successfully");
    Ok(())
}

pub fn cmd_get_input_file(year: usize, day: usize, output: &str) -> anyhow::Result<()> {
    let client = AocClient::new().context("could not build aoc client for future requests")?;

    debug!("aoc client ok");

    let response = client
        .get_input_file(year, day)
        .context(format!("could not get input file for year: {year} and day: {day:02}."))?
        .error_for_status()
        .context(format!("could not get input file for year: {year} and day: {day:02}. Are you sure the session cookie is correctly set up ?"))?;

    debug!("request ok");

    fs::write(output, response.text().unwrap()).context("could not write input data to file")?;

    info!("input data successfully retrieved and saved to '{output}'");
    Ok(())
}

pub fn cmd_submit_answer(year: usize, day: usize, part: &Part, answer: &str) -> anyhow::Result<()> {
    let client = AocClient::new().context("could not build aoc client for future requests")?;

    debug!("aoc client ok");

    let response = client
        .post_answer(year, day, part, answer)
        .context(format!(
            "could not submit '{answer}' for year: {year} and day: {day:02} - part {part}"
        ))?
        .error_for_status()
        .context(format!(
            "could not submit '{answer}' for year: {year} and day: {day:02} - part {part}. Are you sure the session cookie is correctly set up ?"
        ))?;

    debug!("request ok");

    let response_text = response.text()?;

    // Validate answer
    if response_text.contains("That's the right answer!") {
        info!("Correct ! That's the right answer.");
    } else {
        info!("That is not the right answer...");
    }
    Ok(())
}
