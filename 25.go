package main

import (
	"fmt"
	"time"
)

func sleep(d time.Duration) {
	fmt.Println(<-time.After(d)) //after ждет пока пройдет наше время и тогда отправит в канал текущее время
}

func main() {
	fmt.Println("начат отсчет")
	sleep(5 * time.Second)
	fmt.Println("прошло 5 секунд")
}
