package main

import "fmt"

type pessoa struct {
	nome  string
	idade int8
}

type estudante struct {
	pessoa
	curso string
}

func main() {

	var est estudante
	est.pessoa = pessoa{"Reinaldo", 30}
	est.curso = "Engenharia"
	fmt.Println(est)
	fmt.Println(est.pessoa)
	fmt.Println(est.curso)

	p1 := pessoa{"Reinaldo", 30}
	e1 := estudante{p1, "Engenharia"}
	fmt.Println(e1)

}
