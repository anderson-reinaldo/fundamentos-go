package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")

	var array1 [5]string

	array1[0] = "Posicao 0"
	array1[1] = "Posicao 1"
	array1[2] = "Posicao 2"
	array1[3] = "Posicao 3"
	array1[4] = "Posicao 4"

	fmt.Println(array1)

	//array2[5] = "Posicao 5" //Erro pois o array tem 5 posicoes
	array2 := [5]string{"Posicao 0", "Posicao 1", "Posicao 2", "Posicao 3", "Posicao 4"}
	fmt.Println(array2)

	//array3[5] = "Posicao 5" //Erro pois o array tem 5 posicoes
	array3 := [...]string{"Posicao 0", "Posicao 1", "Posicao 2", "Posicao 3", "Posicao 4"}
	fmt.Println(array3)

	//slice é diferente do array pois ele nao tem tamanho fixo.
	slice := []string{"Posicao 0", "Posicao 1", "Posicao 2", "Posicao 3", "Posicao 4"}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	slice = append(slice, "Posicao 5")
	fmt.Println(slice)

	slice = append(slice, "Posicao 6", "Posicao 7", "Posicao 8")
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = "Posicao Alterada"
	fmt.Println(slice2)

	//Arrays Internos
	//Parametro 1: Tipo do slice, Parametro 2: tamanho do slice, Parametro 3: capacidade do slice
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice3 = append(slice3, 5)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice3 = append(slice3, 5)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))
}
