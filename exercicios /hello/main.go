package main

import "fmt"

func main() {
	canal := make(chan string)
	go hello("Hello World!", canal)
	hello := <-canal
	fmt.Println(hello)

}

func hello(message string, canal chan string) {
	canal <- message

}
