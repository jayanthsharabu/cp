use std::io;

fn main() {
    let mut input = String::new();
    io::stdin().read_line(&mut input).unwrap();
    let mut res = String::new();
    for c in input.trim().chars() {
        let lc = c.to_ascii_lowercase();
        match lc {
            'a' | 'o' | 'y' | 'e' | 'u' | 'i' => {
                continue;
            }
            _ => {
                res.push('.');
                res.push(lc);
            }
        }
    }
    println!("{}", res);
}
