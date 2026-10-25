package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Setup fast I/O
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	// Ensure all output is written before the program exits
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	var s string
	fmt.Fscan(reader, &s)

	// Check all pairs of adjacent characters
	for i := 0; i < n-1; i++ {
		if s[i] != s[i+1] {
			fmt.Fprintln(writer, "YES")
			// Print the 2-character diverse substring
			fmt.Fprintln(writer, s[i:i+2])
			return
		}
	}

	// If all adjacent characters are the same, it's impossible
	fmt.Fprintln(writer, "NO")
}
