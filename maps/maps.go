package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario := map[string]string{
		"nome":      "Reinaldo",
		"sobrenome": "Silva",
	}

	fmt.Println(usuario["nome"])
	fmt.Println(usuario["sobrenome"])

	usuario2 := map[string]string{}
	usuario2["nome"] = "Reinaldo"
	usuario2["sobrenome"] = "Silva"

	fmt.Println(usuario2)

	mapAninhado := map[string]map[string]string{
		"usuario1": {
			"nome":      "Reinaldo",
			"sobrenome": "Silva",
		},
		"usuario2": {
			"nome":      "Reinaldo",
			"sobrenome": "Silva",
		},
	}
	fmt.Println(mapAninhado["usuario"]["nome"])

	//Deletando um elemento do map
	delete(mapAninhado, "usuario1")
	fmt.Println(mapAninhado)

	//Adicionando um elemento ao map
	mapAninhado["usuario1"] = map[string]string{
		"nome":      "Reinaldo",
		"sobrenome": "Silva",
	}
	fmt.Println(mapAninhado)

	//Verificando se um elemento existe no map
	_, ok := mapAninhado["usuario1"]
	fmt.Println(ok)

	//Iterando um map
	for chave, valor := range mapAninhado {
		fmt.Println(chave, valor)
	}

	//Iterando um map com range
	for chave, valor := range mapAninhado {
		fmt.Println(chave, valor)
	}

	//Iterando um map com range e usando o underline para ignorar o valor
	for chave := range mapAninhado {
		fmt.Println(chave)
	}
}
