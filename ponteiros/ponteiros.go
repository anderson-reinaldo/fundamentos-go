package main

//Ponteiros sao uma variavel que guarda o endereco de memoria outra variavel

func main() {
	// Fluxo para ponteiros: Define a variavel -> usa & para pegar o endereco da variavel em memoria -> usa * para pegar o valor que o ponteiro aponta
	var variavel1 int = 100
	var variavel2 *int = &variavel1

	println(variavel1, variavel2)
	println(variavel1, *variavel2) //Dereferenciando o ponteiro

}
