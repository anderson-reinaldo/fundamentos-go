package main

import (
	"fmt"
	"sync"
	"time"
)

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}

func main() {

	var waitGroup = &sync.WaitGroup{}
	waitGroup.Add(2)

	go func() {
		escrever("Ola Mundo")
		waitGroup.Done() //decrementa -1 na quantidade de goroutines no waitGroup
	}()

	go func() {
		escrever("Programando em Go")
		waitGroup.Done() //decrementa -1 na quantidade de goroutines no waitGroup
	}()

	waitGroup.Wait() //espera a conclusao de todas as goroutines
}
