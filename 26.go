package main

import (
	"strings"
)

func main() {
	str := strings.ToLower("aabcd")
	chars := []rune(str)
	var unique bool = true
	for i := 0; i < len(chars); i++ {
		for j := 0; j < len(chars); j++ {
			if j != i && chars[i] == chars[j] {
				unique = false
			}
		}
	}
	println(unique)
}
