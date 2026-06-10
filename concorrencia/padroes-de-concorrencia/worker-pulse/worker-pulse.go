package main

import "fmt"

func main() {
	tarefas := make(chan int, 45)
	resultados := make(chan int, 45)

	for w := 1; w <= 3; w++ {
		go worker(tarefas, resultados)
	}

	for i := 1; i <= 45; i++ {
		tarefas <- i
	}

	close(tarefas)

	for a := 1; a <= 45; a++ {
		resultado := <-resultados
		fmt.Println("Resultado: ", a, " = ", resultado)
	}
}

func worker(tarefas <-chan int, resultados chan<- int) {
	for tarefa := range tarefas {
		resultados <- fibonacci(tarefa)
	}
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
