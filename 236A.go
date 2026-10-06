package main

import "fmt"

func chat(name string) {
	mp := make(map[rune]int)
	for _, s := range name {
		mp[s]++
	}
	if len(mp)%2 == 0 {
		fmt.Println("CHAT WITH HER!")
	} else {
		fmt.Println("IGNORE HIM!")
	}
}

func main() {
	var name string
	fmt.Scan(&name)
	chat(name)
}
