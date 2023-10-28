package main

import "fmt"

func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return 1
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func main() {
	posicao := uint(6)
	fmt.Println(fibonacci(posicao))
}
