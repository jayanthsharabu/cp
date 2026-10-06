package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func num_conv(str []string) []int {
	numl := make([]int, len(str))

	for i, s := range str {
		val, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		numl[i] = val
	}
	return numl
}

func str_conv(arr []int) string {
	var sb strings.Builder
	for i, v := range arr {
		if i > 0 {
			sb.WriteString("+")
		}
		sb.WriteString(strconv.Itoa(v))
	}
	return sb.String()
}

func main() {
	var str string
	fmt.Scan(&str)
	runel := strings.Split(str, "+")
	numl := num_conv(runel)
	sort.Ints(numl)
	strl := str_conv(numl)
	fmt.Println(strl)
}
