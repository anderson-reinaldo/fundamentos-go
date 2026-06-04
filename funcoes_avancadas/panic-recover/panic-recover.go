package main

import "fmt"

//panic é uma funcao que interrompe o fluxo da execucao de um programa
//recover é uma funcao que recupera o fluxo da execucao de um programa

func recuperarExecucao() {
	if r := recover(); r != nil {
		fmt.Println("Execucao recuperada com sucesso")
	}
}

func alunoEstaAprovado(n1, n2 int) bool {
	defer recuperarExecucao()

	media := (n1 + n2) / 2
	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A media eh exatamente 6")
}

func main() {

	alunoAprovado := alunoEstaAprovado(6, 6)
	fmt.Println(alunoAprovado)
	fmt.Println("Execucao continua...")
}
