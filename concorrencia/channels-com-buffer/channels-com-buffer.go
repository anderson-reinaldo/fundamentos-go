package main

import "fmt"

func main() {
	//canal com buffer, quando o buffer estiver cheio ele bloqueia a goroutine
	canal := make(chan string, 2)

	canal <- "Ola"
	canal <- "Mundo"
	//Como o canal so tem 2 espacos ele vai dar deadlock
	canal <- "!"

	fmt.Println(<-canal)
	fmt.Println(<-canal)

}
