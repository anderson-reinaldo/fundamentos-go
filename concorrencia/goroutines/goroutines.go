package main

import (
	"fmt"
	"time"
)

func escrever(texto string) {
	fmt.Println(texto)
	time.Sleep(time.Second)
}

func main() {

	go escrever("Ola mundo")
	escrever("Programando em Go")
}
