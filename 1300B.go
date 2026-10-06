package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var t int
	fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		var n int
		fmt.Fscan(reader, &n)
		ts := 2 * n
		arr := make([]int, ts)
		for j := 0; j < ts; j++ {
			fmt.Fscan(reader, &arr[j])
		}
		slices.Sort(arr)
		res := arr[n] - arr[n-1]
		fmt.Fprintln(writer, res)

	}
}
