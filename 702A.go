package main

import "fmt"

func ml(arr []int, n int) int {
	res := 1
	cur := 1
	if n == 1 {
		return cur
	}
	for i := 1; i < n; i++ {
		if arr[i] > arr[i-1] {
			cur++
			if cur > res {
				res = cur
			}
		} else {
			cur = 1
		}
	}

	return res
}

func main() {
	var n int
	fmt.Scan(&n)
	nums := make([]int, n)
	for i := range n {
		fmt.Scan(&nums[i])
	}
	fmt.Println(ml(nums, n))
}
