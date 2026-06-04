package main

import (
	"errors"
	"fmt"
)

func main() {

	// Tipos de Numeros inteiros do GO: int8, int16, int32, int64 quanto maior o inteiros mais espaco ele ocupa
	// int aceita números positivos e negativos
	// uint aceita apenas números positivos (incluindo zero)
	// Tipos de Numeros reais do GO: float32, float64
	// Tipos de Caracteres: string
	// Tipos booleanos: true, false
	// Aspas duplas "" define uma string
	// Aspas simples '' define um char

	var numero int64 = -1000000000000
	fmt.Println(numero)

	var numero2 uint = 123456
	fmt.Println(numero2)

	//alias
	// rune = int32
	var numero3 rune = 123456
	fmt.Println(numero3)

	//byte = uint8
	var numero4 byte = 123
	fmt.Println(numero4)

	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 123123213131.45
	fmt.Println(numeroReal2)

	//STRING

	var texto string = "Texto"
	fmt.Println(texto)

	texto2 := "Texto 2"
	fmt.Println(texto2)

	// Define char e pega o valor na tabela ASCII
	char := 'A'
	fmt.Println(char)

	// FIM STRING

	//BOOLEAN
	var booleano1 bool = true
	fmt.Println(booleano1)

	booleano2 := false
	fmt.Println(booleano2)
	//FIM BOOLEAN

	// ZERO VALUES
	var variavel1 string
	var variavel2 int
	var variavel3 float64
	var variavel4 bool

	fmt.Println(variavel1)
	fmt.Println(variavel2)
	fmt.Println(variavel3)
	fmt.Println(variavel4)

	// FIM ZERO VALUES

	//ERROR
	var error error = errors.New("Erro interno")
	fmt.Println(error)
}
