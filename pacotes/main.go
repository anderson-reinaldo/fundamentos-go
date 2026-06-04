package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo algo na tela no pacote main")
	auxiliar.Escrever1()

	err := checkmail.ValidateFormat("reinaldo@gmail.com")

	if err != nil {
		fmt.Println("Email inválido")
	} else {
		fmt.Println("Email válido")
	}
}
