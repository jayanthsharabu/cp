package main

import "fmt"

func ban(k, n, w int) int {
	tc := k * w * (w + 1) / 2
	if tc <= n {
		return 0
	}
	return tc - n
}

func main() {
	var k, n, w int
	fmt.Scan(&k, &n, &w)
	fmt.Println(ban(k, n, w))
}
