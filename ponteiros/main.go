package main

import "fmt"

func main() {
	var var1 string = "teste ponteiro"
	fmt.Println(var1)
	var2 := &var1
	fmt.Println(*var2)
}
