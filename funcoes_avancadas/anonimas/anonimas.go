package main

import "fmt"

func main() {
	func() {
		fmt.Println("Função Anônima")
	}()

	func(valor string) {
		fmt.Println(valor)
	}("Função Anônima")

	retorno := func(valor string) string {
		return fmt.Sprintf("Retorno: %s", valor)
	}("Teste")
	fmt.Println(retorno)

}
