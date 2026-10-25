use std::io::{self, BufRead};

fn main() {
    let stdin = io::stdin();
    let mut l = stdin.lock().lines();
    if let Some(Ok(fl)) = l.next() {
        let ar: Vec<i64> = fl.split_whitespace().map(|s| s.parse().unwrap()).collect();
        let mut num: i64 = ar[0];
        let k: i64 = ar[1];
        for _ in 0..k {
            if num % 10 == 0 {
                num = num / 10
            } else {
                num = num - 1
            }
        }
        println!("{}", num)
    }
}
