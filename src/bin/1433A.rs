use std::io::{self, BufRead, StdinLock};

fn get_dig(num: u16) -> u16 {
    return num % 10;
}

fn get_digl(num: u16) -> u16 {
    let mut tmp = num;
    let mut ln = 0;
    while tmp > 0 {
        let rem = tmp % 10;
        tmp -= rem;
        tmp /= 10;
        ln += 1;
    }
    return ln;
}
fn get_dsum(num: u16) -> u16 {
    return (num * (num + 1)) / 2;
}
fn key_strokes(num: u16) -> u16 {
    let dg = get_dig(num);
    let dgl = get_digl(num);
    let mut pre = (dg - 1) * 10;
    pre += get_dsum(dgl);
    return pre;
}

fn main() {
    let stdin = io::stdin();
    let mut lines = stdin.lock().lines();
    if let Some(Ok(first_line)) = lines.next() {
        let t: usize = first_line.trim().parse().unwrap_or(0);
        for _ in 0..t {
            if let Some(Ok(line)) = lines.next() {
                let trimmed = line.trim();
                if trimmed.is_empty() {
                    continue;
                }
                let x: u16 = trimmed.parse().unwrap();
                println!("{}", key_strokes(x));
            }
        }
    }
}
