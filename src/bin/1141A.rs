use std::io::{self, BufRead};

fn main() {
    let stdin = io::stdin();
    let mut line = stdin.lock().lines();
    if let Some(Ok(l)) = line.next() {
        let vals: Vec<u64> = l.split_whitespace().map(|s| s.parse().unwrap()).collect();
        let m = vals[0];
        let n = vals[1];

        if n % m != 0 {
            println!("-1");
            return;
        }

        let mut div = n / m;

        let mut count2 = 0;
        while div % 2 == 0 {
            count2 += 1;
            div /= 2;
        }

        let mut count3 = 0;
        while div % 3 == 0 {
            count3 += 1;
            div /= 3;
        }

        if div == 1 {
            println!("{}", count2 + count3);
        } else {
            println!("-1");
        }
    }
}
