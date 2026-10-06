package main

import "fmt"

func remz(ar string) int {
	f, l := 0, 0
	f_flag := false
	for i, v := range ar {
		if v == '1' {
			l = i
			if f_flag == false {
				f = i
				f_flag = true
			}
		}
	}
	if !f_flag {
		return 0
	}
	cnt := 0
	for i := f; i <= l; i++ {
		if ar[i] == '0' {
			cnt++
		}
	}
	return cnt
}

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]string, n)
	for i := range n {
		fmt.Scan(&arr[i])
	}
	for i := range n {
		fmt.Println(remz(arr[i]))
	}
}
