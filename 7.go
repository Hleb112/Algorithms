package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	set := make(map[int]int)
	setMutex := sync.Mutex{}
	var wg sync.WaitGroup
	wg.Add(2)

	for i := 0; i < 20; i++ {
		set[i] = 0
	}

	go func() {
		setMutex.Lock()
		for i := 0; i < 20; i++ {
			set[i] = rand.Intn(100)
		}
		wg.Done()
		setMutex.Unlock()
	}()

	go func() {
		setMutex.Lock()
		for i := 0; i < 20; i++ {
			set[i] = rand.Intn(100)
		}
		wg.Done()
		setMutex.Unlock()
	}()

	wg.Wait()
	fmt.Println(set)
}
