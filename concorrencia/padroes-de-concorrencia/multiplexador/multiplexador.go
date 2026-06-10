package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	canal := multiplexar(escrever("Ola Mundo"), escrever("Programando em Go"))

	for i := 0; i < 5; i++ {
		fmt.Println(<-canal)
	}
}

// multiplexador
func multiplexar(canalDeEntrada1, canalDeEntrada2 <-chan string) <-chan string {
	canalDeSaida := make(chan string)

	go func() {
		for {
			select {
			case msg1 := <-canalDeEntrada1:
				canalDeSaida <- msg1
			case msg2 := <-canalDeEntrada2:
				canalDeSaida <- msg2
			}
		}
	}()
	return canalDeSaida
}

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for i := 0; i < 5; i++ {
			canal <- texto + fmt.Sprintf(" %d", i)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()

	return canal
}
