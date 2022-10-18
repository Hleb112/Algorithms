package main

import (
	"fmt"
	"time"
)

func sleep(d time.Duration) {
	<-time.After(d)
}

func main() {
	fmt.Println("начат отсчет")
	sleep(10 * time.Second)
	fmt.Println("прошло n секунд")
}
