package main

import "fmt"

// func main() {
// 	age := make(map[string][]int)
// 	age["John"] = []int{22, 23, 24, 25}
// 	age["Alice"] = []int{30, 32}
// 	fmt.Println(age["John"][2], age["Alice"][1])
// }

func increaseSalary(salary *int) {
	*salary += 1000
}

func main() {
	employees := make(map[string]*int)

	salaryJhon := 3000
	salaryAlice := 3000

	employees["Jhon"] = &salaryJhon
	employees["Alice"] = &salaryAlice

	//fmt.Println("John's salary:", *employees["Jhon"])
	fmt.Println(employees["Jhon"])

	increaseSalary(employees["Jhon"])

	fmt.Println(employees["Jhon"])
	fmt.Println(employees["Jhon"])
}
