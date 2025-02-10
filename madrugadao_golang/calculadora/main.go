package main

import "fmt"

func main() {
	number := 1288

	soma := 0

	for i := 1; i <= number; i++ {
		soma += i
	}

	fmt.Printf("A soma dos números de 1 a %d é: %d\n", number, soma)
}
