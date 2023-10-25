package main

import "fmt"

func main() {
	var array1 [3]int
	array1[1] = 10
	array1[2] = 10
	fmt.Println(array1)

	slice := []int{10, 11, 12, 13, 14, 15}
	fmt.Println(slice)

	// ARRAY INTERNO
	fmt.Println("-------------------")
	slice2 := make([]int, 10, 15)
	fmt.Println(slice2)
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))
}
