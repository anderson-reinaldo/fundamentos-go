package main

import (
	"fmt"
	"time"
)

// Funcao para inverter o valor do ponteiro
func inverteValor(valor *bool) {
	*valor = !*valor
}

func contador(count *int) {
	*count++
}

func main() {

	var ligado bool = false
	var ponteiro *bool = &ligado
	inverteValor(ponteiro)
	fmt.Println("Valor da variavel: ", ligado)

	for i := 0; i < 10; {
		time.Sleep(time.Second + 1000)
		contador(&i)
		fmt.Println(i)
	}
}
