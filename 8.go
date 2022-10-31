package main

import "fmt"

const dec int64 = 86
const bitNumber = 3

func main() {
	var mask int64 = 1 << bitNumber //создаем маску ..1000 из единицы
	swapOne := dec | mask           //побитовый or поставит 1 где есть хотя бы одна 1
	fmt.Println(swapOne)
	mask = ^mask           //создаем обратную маску ..0111
	swapZero := dec & mask // если где то есть 0 то энд заменит на 0
	fmt.Println(swapZero)
}
