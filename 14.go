package main

import "fmt"

func main() {
	c := make(chan int)

	m := map[int]interface{}{}
	m[1] = 125
	m[2] = "stringg"
	m[3] = true
	m[4] = c

	for key, value := range m {
		switch value.(type) {
		case int:
			fmt.Println(key, "= int")
		case string:
			fmt.Println(key, "= string")
		case bool:
			fmt.Println(key, "= bool")
		case chan int:
			println(key, "= chan")
		default:
			fmt.Println(key, "= undefined")

		}
	}
}
