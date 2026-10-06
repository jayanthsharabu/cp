package main

import "fmt"

func helper(mat []string, n, m int) int {
	res := 0

	for i := 0; i < m-1; i++ {
		if mat[n-1][i] == 'D' {
			res++
		}
	}

	for j := 0; j < n-1; j++ {
		if mat[j][m-1] == 'R' {
			res++
		}
	}

	return res
}

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var n, m int

		fmt.Scan(&n, &m)

		mat := make([]string, n)
		for r := 0; r < n; r++ {
			fmt.Scan(&mat[r])
		}

		fmt.Println(helper(mat, n, m))
	}
}
