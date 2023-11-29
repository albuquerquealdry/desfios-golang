package main

func main() {
	tarefas := make(chan int, 45)
}

func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return 1
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}
