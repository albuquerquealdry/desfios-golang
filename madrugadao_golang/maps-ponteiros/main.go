package main

import "fmt"

func IncriseAuthor(books map[string]string, nome string) {
	books[nome] = "Jostein Gaarder"

}

func main() {
	books := make(map[string]string)

	IncriseAuthor(books, "SofiaWorld")

	fmt.Println(books["SofiaWorld"])

}
