package main

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	posicao := 10
	for i := 0; i <= posicao; i++ {
		println(fibonacci(i))
	}

}
