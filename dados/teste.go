package main

import "fmt"

func main() {
	teste(10 + "ccas")
}

func teste(numero interface{}, texto interface{}) (numero int, texto string) {
	fmt.Println(numero)
	return numero, texto
}
