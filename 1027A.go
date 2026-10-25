package main

import (
	"fmt"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func chck_rev_st(st string) bool {
	rns := []rune(st)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

		diff := abs(int(rns[i]) - int(rns[j]))

		if diff != 0 && diff != 2 {
			return false
		}
	}
	return true
}

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var n int
		var st string
		fmt.Scan(&n, &st)

		if chck_rev_st(st) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
