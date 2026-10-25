package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var n int
	fmt.Fscanln(reader, &n)
	st, _ := reader.ReadString('\n')
	st = strings.TrimSpace(st)
	arr := strings.Split(st, " ")
	mp := make([]int, len(arr))
	for i, v := range arr {
		num, _ := strconv.Atoi(v)
		mp[i] = num
	}
	res := 0

outer:
	for n > 0 {
		for i, v := range mp {
			n -= v
			if n <= 0 {
				res = i + 1
				break outer
			}
		}
	}
	fmt.Fprintln(writer, res)
}
