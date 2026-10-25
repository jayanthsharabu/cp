package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var n int
	var arr string
	fmt.Fscan(reader, &n)
	fmt.Fscan(reader, &arr)
	var res int64 = 0
	for i := 0; i < n; i++ {
		if int(arr[i])%2 == 0 {
			res += int64(i + 1)
		}
	}
	fmt.Fprintln(writer, res)

}
