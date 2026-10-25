use std::io::{self, BufRead};

fn main() {
    let stdin = io::stdin();
    let mut lines = stdin.lock().lines();
    if let Some(Ok(first_line)) = lines.next() {
        let mut nums: Vec<i32> = first_line.split("+").map(|s| s.parse().unwrap()).collect();
        nums.sort();
        let res: String = nums
            .iter()
            .map(|s| s.to_string())
            .collect::<Vec<String>>()
            .join("+");
        println!("{}", res);
    }
}
