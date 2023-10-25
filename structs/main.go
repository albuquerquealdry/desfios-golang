package main

import "fmt"

type usuario struct {
	name  string
	idade uint8
}

func main() {
	var u usuario
	u.name = "aldry"
	u.idade = 21
	fmt.Println(u)
}
