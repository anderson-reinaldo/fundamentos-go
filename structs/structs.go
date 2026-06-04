package main

import "fmt"

type pessoa struct {
	nome     string
	idade    int8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     int8
}

func main() {

	var u pessoa
	u.nome = "Reinaldo"
	u.idade = 30
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua 1", 123}
	u.endereco = enderecoExemplo
	fmt.Println(u)

	u2 := pessoa{"Reinaldo", 30, endereco{"Rua 1", 123}}
	fmt.Println(u2)

	u3 := pessoa{idade: 30}
	fmt.Println(u3)

}
