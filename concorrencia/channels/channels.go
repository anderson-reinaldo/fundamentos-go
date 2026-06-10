package main

import (
	"fmt"
	"time"
)

// OBS se nao fechar o channel ele vai dar deadlock pois ele vai esperar uma resposta e nunca vai receber

func main() {
	channel := make(chan string)

	go escrever("Ola", channel)

	//Aqui so espera a primeira resposta da goroutine
	//msg := <-channel
	//fmt.Println(msg)

	//Jeito 1
	for {
		msg, open := <-channel
		if !open {
			break
		}

		fmt.Println(msg)
	}

	//Jeito 2
	for msg := range channel {
		fmt.Println(msg)
	}

	fmt.Println("Fim do programa")

}

func escrever(text string, ch chan string) {
	for i := 0; i < 5; i++ {
		ch <- text + fmt.Sprintf(" %d", i)
		time.Sleep(time.Second)
	}

	//Fechar o channel para nao dar deadlock
	close(ch)
}
