package main

import (
	"fmt"
	"sync"
)

func main() {
	array := []int{2, 4, 6, 8, 10}
	sum := 0
	var wg sync.WaitGroup
	wg.Add(2)

	go func(array []int) {
		for i := 0; i <= 2; i++ {
			array[i] = array[i] * array[i]

		}
		wg.Done()
	}(array)

	go func(array []int) {
		for i := 3; i <= 4; i++ {
			array[i] = array[i] * array[i]
		}
		wg.Done()
	}(array)
	wg.Wait()

	for i := 0; i < 5; i++ {
		sum = sum + array[i] // сумма квадратов
	}

	fmt.Println(array, sum)
}
