package main

import "fmt"

func changeSalary(salary *int) {
	*salary += 500
}

func main() {
	employees := make(map[string]*int)

	salaryJhon := 3000
	salaryAlice := 3500

	employees["Jhon"] = &salaryJhon
	employees["Alice"] = &salaryAlice

	fmt.Println("Salario Antes: %s, %s", *employees["Jhon"], *employees["Alice"])

	changeSalary(employees["Jhon"])
	fmt.Println("Salario Antes: %s, %s", *employees["Jhon"], *employees["Alice"])

}
