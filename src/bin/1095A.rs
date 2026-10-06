use std::io::{self, Read};

fn cipher(st: &str, n: usize) -> String {
    let mut res = String::with_capacity(n);
    let mut idx: usize = 0;
    let mut step: usize = 1;
    let bytes = st.as_bytes();

    while idx < n {
        let chr: char = bytes[idx] as char;
        res.push(chr);
        idx += step;
        step += 1;
    }

    return res;
}

fn main() {
    let mut input_n = String::new();
    io::stdin().read_line(&mut input_n).unwrap();
    let n: usize = input_n.trim().parse().expect("invalid");
    let mut t = String::new();
    io::stdin().read_to_string(&mut t).unwrap();
    let t = t.trim();
    let dec = cipher(t, n);
    println!("{}", dec);
}
