package main

import (
	"fmt"
	"strings"
)

func rev(str string) string {
	words := strings.Fields(str)
	var res strings.Builder

	//будем записывать в обратном порядке
	for i := len(words) - 1; i >= 0; i-- {
		res.WriteString(words[i])
		res.WriteString(" ")
	}

	return strings.TrimSpace(res.String())
}

func main() {
	str := "snow dog fire cat water mouse"

	strRev := rev(str)
	//fmt.Println(str)
	fmt.Println(strRev)
}
