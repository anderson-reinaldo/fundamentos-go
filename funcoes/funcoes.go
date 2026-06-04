package main

func soma(a int8, b int8) int8 {
	return a + b
}

func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {

	soma := soma(2, 2)
	println(soma)

	soma, subtracao := calculosMatematicos(2, 2)
	println(soma, subtracao)

	// Caso nao queira usar um dos valores retornados use o _ "underline"
	soma1, _ := calculosMatematicos(2, 2)
	println(soma1)

}
