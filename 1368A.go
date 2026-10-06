package main

import "fmt"

func req(a, b, c int) int {
	res := 0
	for a <= c && b <= c {
		if a < b {
			a += b
		} else {
			b += a
		}
		res++
	}
	return res
}

func main() {
	var n int
	fmt.Scan(&n)
	for _ = range n {
		var a, b, n int
		fmt.Scan(&a, &b, &n)
		fmt.Println(req(a, b, n))
	}

}
