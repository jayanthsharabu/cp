package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	flag := false
	for i := 0; i < n; i++ {
		var v int
		fmt.Scan(&v)
		if v%2 == 0 {
			fmt.Println(v / 2)
		} else {
			vn := float64(v) / 2.0
			if flag {
				fmt.Println(int(math.Floor(vn)))
			} else {
				fmt.Println(int(math.Ceil(vn)))
			}
			flag = !flag
		}
	}
}
