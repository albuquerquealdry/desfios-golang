package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)
	go escrever("teste 2", canal)
	for {
		message, open := <-canal
		if !open {
			break
		}
		fmt.Println(message)
	}
	fmt.Println("Fim do loop")

}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto
		time.Sleep(time.Second)
	}
	close(canal)
}
