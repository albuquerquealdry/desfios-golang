package main

import "fmt"

func calc(numeros ...int) int {
	total := 0
	for _, v := range numeros {
		total += v
	}

	return total

}

func main() {
	var soma int = calc(90, 1)
	fmt.Println(soma)
}
