package main

import "fmt"

func main() {
	var palavra string
	fmt.Println("Digite uma palavra")
	fmt.Scan(&palavra)
	fmt.Println(palavra)
}

func sopa(palavra string) string {
	txt := []rune(palavra)
	ret := []rune{}

	/* O -2 é porque começa em 0, e para não pegar a quebra de linha */
	for i := len(txt) - 2; i >= 0; i-- {
		ret = append(ret, txt[i])
	}

	return string(ret)
}
