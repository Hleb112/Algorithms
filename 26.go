package main

import (
	"strings"
)

func main() {
	str := strings.ToLower("aAAAbcd") //понизит регистр
	chars := []rune(str)              // лучше руны тк каждый язык имеет свой размер в байтах
	var unique bool = true
	for i := 0; i < len(chars); i++ {
		for j := 0; j < len(chars); j++ {
			if j != i && chars[i] == chars[j] { //если элементы разные && занчения одинаковые
				unique = false
			}
		}
	}
	println(unique)
}
