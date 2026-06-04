package main

import (
	"fmt"
	"time"
)

func diaDaSemana(dia int) string {
	switch dia {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda"
	case 3:
		return "Terca"
	case 4:
		return "Quarta"
	case 5:
		return "Quinta"
	case 6:
		return "Sexta"
	case 7:
		return "Sabado"
	default:
		return "Dia invalido"
	}
}

func diaDaSemanaComMap(dia int) string {
	diasDaSemana := map[int]string{
		1: "Domingo",
		2: "Segunda",
		3: "Terca",
		4: "Quarta",
		5: "Quinta",
		6: "Sexta",
		7: "Sabado",
	}
	return diasDaSemana[dia]
}

func diaDaSemanaComSlice(dia int) string {
	diasDaSemana := []string{
		"Domingo",
		"Segunda",
		"Terca",
		"Quarta",
		"Quinta",
		"Sexta",
		"Sabado",
	}
	return diasDaSemana[dia-1]
}

func diaDaSemana2(dia int) string {
	switch {
	case dia == 1:
		return "Domingo"
	case dia == 2:
		return "Segunda"
	case dia == 3:
		return "Terca"
	case dia == 4:
		return "Quarta"
	case dia == 5:
		return "Quinta"
	case dia == 6:
		return "Sexta"
	case dia == 7:
		return "Sabado"
	default:
		return "Dia invalido"
	}
}

func diaDaSemanaComFallthrough(dia int) string {
	var diaDaSemana string
	switch {
	case dia == 1:
		diaDaSemana = "Domingo"
	case dia == 2:
		diaDaSemana = "Segunda"
	case dia == 3:
		diaDaSemana = "Terca"
	case dia == 4:
		diaDaSemana = "Quarta"
	case dia == 5:
		diaDaSemana = "Quinta"
		fallthrough
	case dia == 6:
		diaDaSemana = "Sexta"
	case dia == 7:
		diaDaSemana = "Sabado"
	default:
		diaDaSemana = "Dia invalido"
	}

	return diaDaSemana
}

func main() {
	diaNumber := time.Now().Day() + 1
	diaDaSemana(diaNumber)

	fmt.Println(diaDaSemana(diaNumber))
	fmt.Println(diaDaSemana2(diaNumber))
	fmt.Println(diaDaSemanaComMap(diaNumber))
	fmt.Println(diaDaSemanaComSlice(diaNumber))
	fmt.Println("------")
	fmt.Println(diaDaSemanaComFallthrough(diaNumber))

}
