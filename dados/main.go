package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int16 = 100 // pode ser int16, int32 ou int64 depende do tamanho do numero
	fmt.Println(numero)

	var numero2 uint = 1000 // serve para numeros sempre positivos sem sinais de negativo.
	fmt.Println(numero2)

	var numero3 rune = 12345 // rune eh um  alias para o int 32
	fmt.Println(numero3)

	var numero4 byte = 123 // bytre eh igual a uint8
	fmt.Println(numero4)

	var numeroReal1 float32 = 123.32
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 1232134324123321.32
	fmt.Println(numeroReal2)
	// STRINGS

	var str string = "Texto"
	fmt.Println(str)

	char := 'A' // CHAR SEMPRE VAI RETORNAR A REFERENCIA DO CARACTWERE NA TABELA ASSIC
	fmt.Println(char)

	// BOLEANOS

	var booleano bool = true
	fmt.Println(booleano)

	var erro error
	fmt.Println(erro)

	var meuErro error = errors.New("nascer")
	fmt.Println(meuErro)
}
