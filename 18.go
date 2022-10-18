package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type S struct {
	counter atomic.Int64
}

func (s *S) SetCounter(v int64) {
	s.counter.Store(v)
}

func (s *S) GetCounter() int64 {
	return s.counter.Load()
}

func (s *S) Add(v int64) {
	s.counter.Add(v)
}

func main() {
	var counter S

	counter.SetCounter(0)
	for i := 0; i < 1000; i++ {
		go counter.Add(1)
	}
	time.Sleep(1)
	fmt.Println(counter.GetCounter())
}
