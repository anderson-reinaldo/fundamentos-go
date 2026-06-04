package main

func main() {

	numero := 1

	if numero > 2 {
		println("Isso é verdadeiro")
	} else {
		println("Isso nao é verdadeiro")
	}

	// if init; condicao; else { }, as variaveis declaradas dentro do if nao podem ser usadas fora dele
	if outroNumero := numero; outroNumero > 0 {
		println("Esee numero é maior que zero")
	} else if outroNumero < 0 {
		println("Esee numero é menor que zero")
	} else {
		println("Esee numero é zero")
	}

}
