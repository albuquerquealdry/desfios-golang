package main

import (
	"fmt"
	"time"
)

func main() {
	canal1, canal2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			canal1 <- "Canal 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Millisecond * 2)
			canal2 <- "Canal 2"
		}
	}()

	for {

		select {
		case messageCanal1 := <-canal1:
			fmt.Println(messageCanal1)
		case messageCanal2 := <-canal2:
			fmt.Println(messageCanal2)
		}

	}

}
