package main

import "fmt"

func main() {
	i := 0
	for i < 10 {
		i++
		fmt.Println("Incrmentando i")
		time.Sleep(time.Second)
	}

	j := 0

	name := [3]string{"a", "b", "c"}
	for i, v := range name {
		fmt.Println(i, v)
	}

	name := []string{"a", "b", "c"}
	teste2 := &name
	fmt.Println(teste2)
	for _, v := range name {
		teste := &v
		fmt.Println(*teste)

	}
}
