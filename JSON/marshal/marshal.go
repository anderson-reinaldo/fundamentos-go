package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome string `json:"nome"`
	Raca string `json:"raca"`
}

func main() {
	// json.Marshal
	bob := cachorro{"Bob dos ovão bem grandão", "Dalmata"}
	c, err := json.Marshal(bob)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(bytes.NewBuffer(c))

	cachorro2 := map[string]string{
		"nome": "Bob dos ovão bem grandão",
		"raca": "Dalmata",
	}

	c2, err := json.Marshal(cachorro2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(c2))
}
