use std::{
    cmp,
    io::{self, BufRead},
};

fn lookl(arr: &[usize], x: usize, i: usize) -> bool {
    let val = arr[i];
    let im = i.saturating_sub(x);
    for j in im..i {
        if arr[j] <= val {
            return false;
        }
    }
    true
}

fn lookr(arr: &[usize], y: usize, i: usize) -> bool {
    let val = arr[i];
    let im = cmp::min(y + i, arr.len() - 1);
    for j in i + 1..=im {
        if arr[j] <= val {
            return false;
        }
    }
    true
}

fn fday(arr: &[usize], x: usize, y: usize, n: usize) -> usize {
    for i in 0..n {
        if lookl(arr, x, i) && lookr(arr, y, i) {
            return i + 1;
        }
    }
    0
}

fn main() {
    let stdin = io::stdin();
    let mut lines = stdin.lock().lines();
    if let Some(Ok(first_line)) = lines.next() {
        let parts: Vec<usize> = first_line
            .split_whitespace()
            .filter_map(|s| s.parse().ok())
            .collect();
        let n = parts[0];
        let x = parts[1];
        let y = parts[2];
        if let Some(Ok(second_line)) = lines.next() {
            let arr: Vec<usize> = second_line
                .split_whitespace()
                .filter_map(|s| s.parse().ok())
                .collect();
            let res = fday(&arr, x, y, n);
            println!("{}", res);
        }
    }
}
