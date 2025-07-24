use std::{env, fs};

use reqwest::{
    blocking::{self, Response},
    header,
};

use crate::{
    Part,
    cookies::session::{CONFIG_DIR, SESSION_FILE},
};

const BASE_URL: &str = "https://adventofcode.com";

pub struct AocClient {
    client: blocking::Client,
    base_url: String,
}

impl AocClient {
    pub fn new() -> Result<AocClient, reqwest::Error> {
        let home_dir = env::home_dir().expect("cannot read HOME var env");
        let session = fs::read_to_string(home_dir.join(CONFIG_DIR).join(SESSION_FILE))
            .expect("missing session file. Did you correctly set up the session using get-session or set-session commands ?");

        let mut headers = header::HeaderMap::new();

        let mut session_value =
            header::HeaderValue::from_str(format!("session={session}").as_str()).unwrap();
        session_value.set_sensitive(true);

        headers.insert(header::COOKIE, session_value);

        Ok(AocClient {
            client: blocking::Client::builder()
                .default_headers(headers)
                .build()?,
            base_url: String::from(BASE_URL),
        })
    }

    pub fn get_input_file(&self, year: usize, day: usize) -> Result<Response, reqwest::Error> {
        let url = format!("{}/{year}/day/{day}/input", self.base_url);
        self.client.get(url).send()
    }

    pub fn post_answer(
        &self,
        year: usize,
        day: usize,
        part: &Part,
        answer: &str,
    ) -> Result<Response, reqwest::Error> {
        let url = format!("{}/{year}/day/{day}/answer", self.base_url);

        let part_str = match part {
            Part::One => "1",
            Part::Two => "2",
        };

        self.client
            .post(url)
            .form(&[("level", part_str), ("answer", answer)])
            .send()
    }
}
