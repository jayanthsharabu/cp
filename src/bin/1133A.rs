use std::io::{self, BufRead};

fn main() {
    let stdin = io::stdin();
    let mut lines = stdin.lock().lines();
    if let Some(Ok(fl)) = lines.next() {
        if let Some(Ok(sl)) = lines.next() {
            let a1: Vec<i32> = fl.split(":").map(|s| s.parse::<i32>().unwrap()).collect();
            let a2: Vec<i32> = sl.split(":").map(|s| s.parse::<i32>().unwrap()).collect();
            let h1: i32 = a1[0];
            let m1: i32 = a1[1];
            let h2: i32 = a2[0];
            let m2: i32 = a2[1];
            let mid_min: i32 = ((h2 * 60 + m2) + (h1 * 60 + m1)) / 2;
            let rm: i32 = mid_min - (mid_min % 60);
            println!("{:02}:{:02}", (rm / 60), (mid_min % 60))
        }
    }
}
