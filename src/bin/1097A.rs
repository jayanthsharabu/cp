use std::io::{self, Read};

fn check(tc: &str, tcl: &[String]) -> bool {
    let tcb = tc.as_bytes();
    let rnk = tcb[0];
    let st = tcb[1];

    for card in tcl {
        let cb = card.as_bytes();
        if cb[0] == rnk || cb[1] == st {
            return true;
        }
    }
    return false;
}

fn main() {
    let mut input = String::new();
    io::stdin().read_to_string(&mut input).unwrap();
    let mut tokens = input.split_whitespace();
    let tc = tokens.next().unwrap();
    let tcl: Vec<String> = tokens.take(5).map(|s| s.to_string()).collect();
    if check(tc, &tcl) {
        println!("YES");
    } else {
        println!("NO")
    }
}
