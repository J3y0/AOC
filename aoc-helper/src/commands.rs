use std::{
    env, fs,
    path::{Path, PathBuf},
};

use crate::cookies::session;
use crate::{Part, client::AocClient};

pub fn cmd_set_session(session: &str) {
    let home_path = env::home_dir().expect("cannot get HOME var env");
    let config_dir_path = home_path.join(".config/aoc-helper/");
    if !fs::exists(&config_dir_path).unwrap() {
        match fs::create_dir(&config_dir_path) {
            Ok(_) => {}
            Err(e) => eprintln!("error creating config dir: {}", e),
        }
    }

    match fs::write(config_dir_path.join("session"), session) {
        Ok(_) => println!("session set up successfully"),
        Err(e) => eprintln!("error writting to session file: {}", e),
    }
}

pub fn cmd_get_session() {
    let home_path = env::home_dir().expect("cannot get HOME varenv");
    let search_path = home_path.join(".mozilla/firefox/*.default-release/cookies.sqlite");

    let mut path = PathBuf::new();
    for entry in glob::glob(search_path.to_str().unwrap()).expect("failed to read glob pattern") {
        if let Ok(ent) = entry {
            path = ent;
        }
    }

    // tmp database if firefox open
    let tmp_path = Path::new("/tmp").join(path.file_name().unwrap());
    println!("tmp_path: {:?}", &tmp_path);
    let _ = fs::copy(path, &tmp_path);

    match session::retrieve_session(&tmp_path) {
        Ok(session) => {
            println!("session retrieved successfully: \"{}\"", session);
            cmd_set_session(&session);
        }
        Err(e) => eprintln!("error retrieving session: {}", e),
    }

    fs::remove_file(tmp_path).expect("cannot remove tmp database file");
}

pub fn cmd_get_input_file(year: usize, day: usize, output: &str) {
    let client = match AocClient::new() {
        Ok(c) => c,
        Err(e) => panic!("error building AoC client: {}", e),
    };

    let resp = match client.get_input_file(year, day) {
        Ok(res) => res,
        Err(e) => panic!("reqwest error: {}", e),
    };

    fs::write(output, resp).expect("cannot write to file");
}

pub fn cmd_submit_answer(year: usize, day: usize, part: Part, answer: &str) {
    let client = match AocClient::new() {
        Ok(c) => c,
        Err(e) => panic!("error building AoC client: {}", e),
    };

    let resp = match client.post_answer(year, day, part, answer) {
        Ok(res) => res,
        Err(e) => panic!("reqwest error: {}", e),
    };

    // Validate answer
    if resp.contains("That's the right answer!") {
        println!("Correct ! That's the right answer.");
    } else {
        println!("That is not the right answer...");
    }
}
