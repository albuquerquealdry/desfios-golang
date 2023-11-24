package main

import (
	"fmt"
	"time"
)

func main() {
	go escrever("teste")
	escrever("teste2")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
