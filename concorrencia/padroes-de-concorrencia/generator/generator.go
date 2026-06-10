package main

import (
	"fmt"
	"time"
)

func main() {
	canal := escrever("Ola Mundo")
	for i := 0; i < 5; i++ {
		fmt.Println(<-canal)
	}
}

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for i := 0; i < 5; i++ {
			canal <- texto + fmt.Sprintf(" %d", i)
			time.Sleep(time.Second)
		}
	}()

	return canal
}
