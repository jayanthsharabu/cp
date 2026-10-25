package main

import "fmt"

func c1(a, b, c int) int {
	if c > a {
		return -1
	}
}

func c2(a, b, c int) int {}

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var a, b, c int
		fmt.Scan(&a, &b, &c)
		fmt.Println(c1(a, b, c), c2(a, b, c))
	}
}
