use std::io::{self, BufRead};

fn one_chck(st: String) -> String {
    for v in st.chars() {
        if v == '1' {
            return String::from("HARD");
        }
    }
    return String::from("EASY");
}

fn main() {
    let stdin = io::stdin();
    let mut lines = stdin.lock().lines();
    if let Some(Ok(l)) = lines.next() {
        let t: i32 = l.parse().unwrap();
        if let Some(Ok(l2)) = lines.next() {
            let fl: String = l2.split_whitespace().collect();
            println!("{}", one_chck(fl));
        }
    }
}
