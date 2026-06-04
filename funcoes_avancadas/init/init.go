package main

var n int

// init e executado quando o pacote e carregado sempre antes do main e todo aquivo pode ter uma funcao init diferente da funcao main que é uma por pacote
func init() {
	n = 10
	println("Executando init")
}

func main() {
	println("Executando main e o valor de n é:", n)
}
