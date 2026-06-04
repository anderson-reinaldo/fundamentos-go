package main

func main() {
	// ARITHMETIC OPERATORS
	soma := 1 + 1
	subtracao := 1 - 1
	multiplicacao := 1 * 1
	divisao := 1 / 1
	restoDaDivisao := 1 % 1

	println(soma, subtracao, multiplicacao, divisao, restoDaDivisao)

	// o GO nao pemite somar, dividir ou fazer qualquer operação com variaveis de tipos diferentes ex: int e float ou int8 e int32
	var numero1 int16 = 10
	var numero2 int16 = 5
	soma2 := numero1 + numero2
	println(soma2)
	// FIM DE ARITHMETIC OPERATORS

	//ATRIBUIÇÃO
	var variavel1 string = "String"
	variavel2 := "String 2"
	println(variavel1, variavel2)
	// FIM DE ATRIBUIÇÃO

	//OPERADORES RELACIONAIS
	igual := 1 == 1
	diferente := 1 != 1
	maior := 1 > 1
	menor := 1 < 1
	maiorOuIgual := 1 >= 1
	menorOuIgual := 1 <= 1
	println(igual, diferente, maior, menor, maiorOuIgual, menorOuIgual)
	// FIM DE OPERADORES RELACIONAIS

	//OPERADORES LOGICOS
	verdadeiro, falso := true, false
	println(verdadeiro && falso)
	println(verdadeiro || falso)
	println(!verdadeiro)
	// FIM DE OPERADORES LOGICOS

	//OPERADORES UNÁRIOS
	count := 10
	count++
	count--
	count += 10
	count -= 10
	count *= 10
	count /= 10
	count %= 10
	println(count)
	// FIM DE OPERADORES UNÁRIOS

	//OPERADORES TERNARIOS
	var idade = 20
	var mensagem string
	if idade >= 18 {
		mensagem = "Maior de idade"
	} else {
		mensagem = "Menor de idade"
	}
	println(mensagem)
	// FIM DE OPERADORES TERNARIOS
}
