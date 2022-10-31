package main

import (
	"fmt"
	"time"
)

func worker(jobs <-chan int, results chan<- int, id int) { //jobs только получение, res только отправка
	for j := range jobs { //перебираем значения из канала
		time.Sleep(time.Millisecond) // simulating blocking task
		fmt.Printf("[worker %v] Sending result by worker %v\n", id, id)
		results <- j * 2
	}
}

func main() {
	jobs := make(chan int, 50) // 50 - размер канала буфер
	results := make(chan int, 50)

	for i := 0; i < 3; i++ {
		go worker(jobs, results, i) //создали 3 рутины
	}

	for j := 0; j < 20; j++ { //отправили 20 работ
		jobs <- j //non-blocking as buffer capacity is 50
	}

	close(jobs)

	for i := 0; i < 20; i++ {
		result := <-results //блок тк буфер пуст, ждем пока воркеры вернут результат
		fmt.Println(result)
	}
}
