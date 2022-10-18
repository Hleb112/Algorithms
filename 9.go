package main

import "fmt"

/*
Генерируем данные из слайса
отправляем и в канал
*/
func gen1(data []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, i := range data {
			out <- i
		}
		close(out)
	}()
	return out
}

/*
Читаем данных из канала
обрабатываем их
отправляем в канал
*/
func work(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for i := range in {
			out <- i * 2
		}
		close(out)
	}()
	return out
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}

	for res := range work(gen1(arr)) {
		fmt.Println(res)
	}
}
