package main

import (
	"fmt"
	"sync"
)

func main() {
	set := make(map[int]int)
	setMutex := sync.Mutex{}
	var wg sync.WaitGroup
	wg.Add(2)

	for i := 0; i < 20; i++ { //забиваем карту нулями
		set[i] = 0
	}

	go func() {
		setMutex.Lock() //закрыть мьютекс перед доступом к мап
		for i := 0; i < 20; i++ {
			set[i] = 1 //rand.Intn(100)
		}
		wg.Done() //горутина закончила
		fmt.Println("рутина 1")
		setMutex.Unlock() //открыть мьютекс
	}()

	go func() {
		setMutex.Lock()
		for i := 0; i < 20; i++ {
			set[i] = 2 //rand.Intn(100)
		}
		wg.Done()
		fmt.Println("рутина 2")
		fmt.Println(set)
		setMutex.Unlock()
	}()

	wg.Wait()
	fmt.Println(set)
}
