package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"time"
)

func setParam() (string, error) {
	switch len(os.Args) {
	case 1:
		return "", errors.New("too few arguments... plz add second argument")

	case 2:
		return os.Args[1], nil //len 2 потому что [0] это путь

	default:
		return "", errors.New("too many arguments")
	}
}

func read1(data <-chan int) {
	for i := range data {
		fmt.Println(i)
	}

}

func main() {

	param, err := setParam()
	if err != nil {
		log.Fatalln(err)
	}

	t, err := strconv.Atoi(param) // t- аргумент вписанный (секунды) / Atoi - преобразует param из стринг в int
	if err != nil {
		fmt.Println(err)
	}

	data := make(chan int)
	timeout := time.After(time.Duration(t) * time.Second) //после таймаута отправляет в канал информацию время t

	go read1(data)

	go func() {
		sigchan := make(chan os.Signal)
		signal.Notify(sigchan, os.Interrupt)
		<-sigchan
		log.Println("Program killed !")

		// do last actions and wait for all write operations to end

		os.Exit(0)
	}()

	var i int
	for {
		// select блокируется до тех пор, пока один из блоков case не будет выполнен
		select {
		case <-timeout:
			fmt.Println("timeout")
			close(data)
			return

		case data <- rand.Intn(10):
			fmt.Println(i)

		}
	}
}
