package main

import "fmt"

type Speaker interface {
	Speak() string
}

type Person struct {
	Name string
}

func (p Person) Speak() string {
	return "Hello, my name is " + p.Name
}

type Dog struct {
	Breed string
}

func (d Dog) Speak() string {
	return "Au Au, I am a " + d.Breed
}

func main() {
	p := Person{Name: "João"}
	fmt.Println(p.Speak())

	d := Dog{Breed: "Labrador"}
	fmt.Println(d.Speak())
}
