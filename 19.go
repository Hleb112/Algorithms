package main

import "fmt"

//функция берет стринг на вход и возвращает перевернутую на выход
func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func main() {
	str := "майнкрафт это моя жизнь"

	strRev := reverse(str)
	fmt.Println(str)
	fmt.Println(strRev)
}
