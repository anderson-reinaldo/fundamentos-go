package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func (p pessoa) nomeCompleto() string {
	return p.nome + " " + p.sobrenome
}

func (p *pessoa) fazerAniversario() {
	p.idade++
}

func main() {
	var pessoa pessoa = pessoa{
		nome:      "Anderson",
		sobrenome: "Reinaldo",
		idade:     25,
	}

	resultMetodo := pessoa.nomeCompleto()
	fmt.Println(resultMetodo)
	pessoa.fazerAniversario()
	fmt.Println(pessoa)
}
