package main

import "fmt"

func closure() func() {
	texto := "Reinaldo"
	return func() {
		fmt.Println(texto)
	}
}

func main() {
	texto := "Iniciando main"
	fmt.Println(texto)

	f := closure()
	f()

}
