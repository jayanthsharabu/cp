package main

import (
	"fmt"
)

func subtract(x int) int {
	rem := x % 10
	if rem != 0 {
		x -= 1
	} else {
		x /= 10
	}
	return x
}

func subtract_n(x, y int) int {
	// Changed to a standard loop for compatibility
	for _ = range y {
		x = subtract(x)
	}
	return x
}

func main() {
	var x, y int
	fmt.Scan(&x, &y)
	fmt.Println(subtract_n(x, y))
}
