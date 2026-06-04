package main

import "fmt"

func media(n1, n2 float64) float64 {
	total := (n1 + n2) / 2
	defer fmt.Println("Media", total)
	if total > 6 {
		fmt.Println("Aprovado", total)
	} else {
		fmt.Println("Reprovado", total)
	}
	return total
}

func main() {
	defer func() {
		fmt.Println("Execultei")
	}()
	println("Executando")
	media(8, 7)
}
