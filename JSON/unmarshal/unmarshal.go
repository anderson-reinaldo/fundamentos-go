package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome string `json:"-"`
	Raca string `json:"raca"`
}

func main() {
	cachorroEmJSON := `{"nome": "Bob", "raca": "Dalmata"}`
	var c cachorro

	// Primeiro parametro manda o JSON convertido para um slice de bytes, Segundo parametro manda o endereço da variavel que vai receber o STRUCT
	//O json.Unmarshal retorna somente um erro
	if err := json.Unmarshal([]byte(cachorroEmJSON), &c); err != nil {
		log.Fatal(err)
	}

	//[]byte(cachorroEmJSON) converte o JSON para um slice de bytes
	fmt.Println([]byte(cachorroEmJSON))
	fmt.Println(c)

	cachorro2 := make(map[string]string)
	if err := json.Unmarshal([]byte(cachorroEmJSON), &cachorro2); err != nil {
		log.Fatal(err)
	}

	fmt.Println(cachorro2)

}
