package main

import "fmt"

func main() {
	note := 7

	switch {
	case note < 7:
		fmt.Println("Recuperacao")
	case note > 7:
		fmt.Println("Passou")
	}

}
