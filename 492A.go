package main

import "fmt"

func req_c(c int) int {
	height := 0
	totalUsed := 0

	for {
		nextLevel := height + 1
		cubesForNextLevel := (nextLevel * (nextLevel + 1)) / 2

		if totalUsed+cubesForNextLevel > c {
			break
		}

		totalUsed += cubesForNextLevel
		height++
	}

	return height
}

func main() {
	var c int
	fmt.Scan(&c)
	fmt.Println(req_c(c))
}
