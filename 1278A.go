package main

import "fmt"

func fc(s1, s2 string) bool {
	n := len(s1)
	var mp1, mp2 [26]int
	for i := 0; i < n; i++ {
		mp1[s1[i]-'a']++
		mp2[s2[i]-'a']++
	}
	return mp1 == mp2
}

func fcc(s1, s2 string) string {
	l := len(s1)
	n := len(s2)
	for i := 0; i <= n-l; i++ {
		if fc(s1, s2[i:i+l]) {
			return "YES"
		}
	}
	return "NO"
}

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var s1, s2 string
		fmt.Scan(&s1)
		fmt.Scan(&s2)
		fmt.Println(fcc(s1, s2))
	}
}
