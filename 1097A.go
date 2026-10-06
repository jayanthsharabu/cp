package main

import "fmt"

func check(tc string, tcl []string) bool {
	rnk := tc[0]
	st := tc[1]

	for i := range 5 {
		if tcl[i][0] == rnk || tcl[i][1] == st {
			return true
		}
	}
	return false
}

func main() {
	var tc string
	fmt.Scan(&tc)
	tcl := make([]string, 5)
	for i := range 5 {
		fmt.Scan(&tcl[i])
	}
	if check(tc, tcl) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
