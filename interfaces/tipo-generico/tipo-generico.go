package main

// Interfaces podem ser usadas para criar tipos genericos GAMBIARRA, porem na versão 1.18 nao é mais usada pois o generics foi introduzido

import "fmt"

func generica(inter interface{}) {
	fmt.Println(inter)
}

func main() {
	generica(1)
	generica("Reinaldo")
	generica(true)

	//onde pode usar mais nao deve ser usado!!!!
	mapa := map[interface{}]interface{}{
		1:            "String",
		"String":     2,
		float32(100): true,
		true:         float32(100),
	}
	fmt.Println(mapa)
}
