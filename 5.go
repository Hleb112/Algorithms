package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func setParam() (string, error) {
	switch len(os.Args) {
	case 1:
		return "", errors.New("too few arguments... plz add second argument")

	case 2:
		return os.Args[1], nil

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

	t, err := strconv.Atoi(param)
	if err != nil {
		fmt.Println(err)
	}

	data := make(chan int)
	timeout := time.After(time.Duration(t) * time.Second)

	go read1(data)

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
