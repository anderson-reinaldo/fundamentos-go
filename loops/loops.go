package main

import "time"

// No GO só existe um laço, o for.
// Não é possivel usar range em STRUCT

func main() {
	// Maneira 1, igual ao while
	i := 0
	for i < 10 {
		println(i)
		i++
	}

	// Maneira 2, define a variavel, condicional e o incrementador
	for i := 0; i < 10; i++ {
		println(i)
	}

	// Maneira 3, define a variavel, condicional e o incrementador
	for i := 0; i < 10; {
		println(i)
		i++
	}

	//Maneira 4: intera no slice ou array OBS: caso nao queira usar o index use o _ "underline"
	numeros := [3]string{"um", "dois", "tres"}
	for index, valor := range numeros {
		println(index, valor)
	}

	// Maneira 5: intera no map
	usuario := map[string]string{
		"nome":      "Reinaldo",
		"sobrenome": "Silva",
	}
	for chave, valor := range usuario {
		println(chave, valor)
	}

	// Maneira 6: intera em uma string
	for _, valor := range "Reinaldo" {
		println(string(valor))
	}

	//Loop infinito
	for {
		time.Sleep(time.Second)
		println("Loop infinito")
	}
}
