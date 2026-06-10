package main

import (
	"fmt"
	"time"
)

// o SELECT é uma estrutura de controle que permite executar multiplos fluxos de execução ao mesmo tempo
// ele funciona da mesma maneira que o switch, mas ele permite executar multiplos fluxos de execução ao mesmo tempo

func main() {
	chanel1 := make(chan string)
	channel2 := make(chan string)

	go func() {
		for {
			time.Sleep(time.Second)
			chanel1 <- "Resultado 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			channel2 <- "Resultado 2"
		}
	}()

	for {
		select {
		case res1 := <-chanel1:
			fmt.Println(res1)
		case res2 := <-channel2:
			fmt.Println(res2)
		}
	}

}
