package main

//Não consigo ter mais de um parametro variatico por funcao ex: func soma(numeros ...int, numeros2 ...int) e sempre o variatico deve ser o ultimo

import "fmt"

func soma(numeros ...int) (total int) {
	total = 0
	for _, numero := range numeros {
		total += numero
	}
	return
}

func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	totalSoma := soma(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(totalSoma)
	escrever("Escrevendo", 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
}
