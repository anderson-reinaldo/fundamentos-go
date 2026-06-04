package main

// O que saiba sobre Generics incluidos na versão 1.18 do GO
// Generics sao tipos que podem ser usados para criar tipos genericos, ex: int | float64 um parametro que pode ser int ou float e nao fixos como só int ou so float
import "fmt"

type number interface {
	int | float64
}

// Em uso de Generics incluidos na versão 1.18 do GO
func calculos[T number](n1, n2 T) T {
	return n1 + n2
}

func main() {
	//nesse caso eu nao preciso passa o int ou float na posicao correta o calculo vai ser feito pois eu estou usando Generics que pode ser int ou float em qualquer posicao dos parametros
	soma := calculos(1, 2.5)
	fmt.Println(soma)
}
