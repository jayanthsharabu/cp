use std::io::{self, BufRead};

fn main() {
    let stdin = io::stdin();
    let mut lines = stdin.lock().lines();

    if let Some(Ok(line)) = lines.next() {
        let s = line.trim();

        let is_caps_accident = s.chars().skip(1).all(|c| c.is_ascii_uppercase());

        if is_caps_accident {
            let fixed: String = s
                .chars()
                .map(|c| {
                    if c.is_ascii_lowercase() {
                        c.to_ascii_uppercase()
                    } else {
                        c.to_ascii_lowercase()
                    }
                })
                .collect();

            println!("{}", fixed);
        } else {
            println!("{}", s);
        }
    }
}
