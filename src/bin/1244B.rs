use std::{
    io::{self, BufRead},
    usize,
};

fn rv(n: usize, s: &[u8]) -> usize {
    let mut first: Option<usize> = None;
    let mut last: Option<usize> = None;

    for i in 0..n {
        if s[i] == b'1' {
            if first.is_none() {
                first = Some(i);
            }
            last = Some(i);
        }
    }
    match (first, last) {
        (Some(f), Some(l)) => {
            let ld = 2 * (l + 1);
            let rd = 2 * (n - f);
            n.max(ld).max(rd)
        }
        _ => n,
    }
}

fn main() {
    let stdin = io::stdin();
    let mut lines = stdin.lock().lines();
    if let Some(Ok(first_line)) = lines.next() {
        let t: usize = first_line.trim().parse().unwrap_or(0);
        for _ in 0..t {
            if let Some(Ok(l1)) = lines.next() {
                let n: usize = l1.trim().parse().unwrap_or(0);
                if let Some(Ok(l2)) = lines.next() {
                    let arr = l2.trim().as_bytes();
                    println!("{}", rv(n, arr));
                }
            }
        }
    }
}
