package main

import "fmt"

func bear(a, b int) int {
	res := 0
	for {
		if a > b {
			return res
		} else {
			a *= 3
			b *= 2
			res += 1
		}
	}
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(bear(a, b))
}
