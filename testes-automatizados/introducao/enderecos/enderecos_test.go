package enderecos_test

import (
	. "introducao-a-testes/enderecos"
	"testing"
)

//Teste de unidade

type cenarioDeTeste struct {
	enderecoInserido string
	tipoEsperado     string
}

func TestTipoEndereco(t *testing.T) {
	t.Parallel()

	cenariosDeTeste := []cenarioDeTeste{
		{"Rua", "Rua"},
		{"Avenida", "Avenida"},
		{"Estrada", "Estrada"},
		{"Rodovia", "Rodovia"},
		//{"", "Tipo de endereço inválido"},
		//{"Praça", "Tipo de endereço inválido"},
		{"Rodovia", "Rodovia"},
	}

	for _, cenario := range cenariosDeTeste {
		retornoRecebido := TipoEndereco(cenario.enderecoInserido)
		if retornoRecebido != cenario.tipoEsperado {
			t.Errorf("O tipo de endereço esperado é %s e o recebido é %s", cenario.tipoEsperado, retornoRecebido)
		}
	}
}

func TestQualquer(t *testing.T) {
	t.Parallel()
	if 1 > 2 {
		t.Errorf("1 não é maior que 2")
	}
}
