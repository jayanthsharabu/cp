package main

import (
	"fmt"
	"strings"
)

func arr_print(n, N int) {
	spaces := strings.Repeat("  ", N-n)
	res := []string{"0"}
	for i := 1; i <= n; i++ {
		res = append(res, fmt.Sprint(i))
	}
	for i := n - 1; i > 0; i-- {
		res = append(res, fmt.Sprint(i))
	}
	res = append(res, "0")
	fmt.Print(spaces)
	fmt.Println(strings.Join(res, " "))
}

func rhombus_print(n int) {
	if n < 0 {
		return
	}
	fmt.Print(strings.Repeat("  ", n))
	fmt.Println("0")
	if n > 0 {
		for i := 1; i <= n; i++ {
			arr_print(i, n)
		}

		for i := n - 1; i >= 1; i-- {
			arr_print(i, n)
		}
	}
	fmt.Print(strings.Repeat("  ", n))
	fmt.Println("0")
}

func main() {
	var n int
	fmt.Scan(&n)
	rhombus_print(n)
}
