package main

import "fmt"

func capitalize(name []byte) []byte {
	var chr byte
	chr = name[0]
	//A = 65, a=97, z = 122, Z = 90
	num := int(chr)
	if num <= 122 && num >= 97 {
		name[0] -= 32
	}
	return name

}
func main() {
	var name []byte
	fmt.Scan(&name)
	fmt.Println(string(capitalize(name)))
}
