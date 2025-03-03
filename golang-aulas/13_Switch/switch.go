package main

import (
	"fmt"
)

// Primeira maneira de criar um switch declarando o valor na clausula inicial
func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-Feira"
	case 3:
		return "Terça-Feira"
	case 4:
		return "Quarta-Feira"
	case 5:
		return "Quinta-Feira"
	case 6:
		return "Sexta-Feira"
	case 7:
		return "Sábado"
	default:
		return "Número Inválido"
	}

}

func diaDaSemana2(numero int) string {
	switch {
	case numero == 1:
		return "Domingo"
	case numero == 2:
		return "Segunda-Feira"
	case numero == 3:
		return "Terça-Feira"
	case numero == 4:
		return "Quarta-Feira"
	case numero == 5:
		return "Quinta-Feira"
	case numero == 6:
		return "Sexta-Feira"
	case numero == 7:
		return "Sábado"
	default:
		return "Número Inválido"
	}
}

func diaDaSemana3(numero int) string {
	var diaDaSemana string

	switch {
	case numero == 1:
		diaDaSemana = "Domingo"
		fallthrough
	case numero == 2:
		diaDaSemana = "Segunda-Feira"
	case numero == 3:
		diaDaSemana = "Terça-Feira"
	case numero == 4:
		diaDaSemana = "Quarta-Feira"
	case numero == 5:
		diaDaSemana = "Quinta-Feira"
	case numero == 6:
		diaDaSemana = "Sexta-Feira"
	case numero == 7:
		diaDaSemana = "Sábado"
	default:
		diaDaSemana = "Número Inválido"
	}

	return diaDaSemana
}

// 13 - Aula sobre Switch em GO
func main() {
	fmt.Println("Switch")

	// Chamando a função do diaDaSemana
	dia := diaDaSemana(3)
	fmt.Println("Dia da semana: ", dia)

	// Chamando a função do diaDaSemana para função 2
	dia2 := diaDaSemana2(6)
	fmt.Println("Dia da semana função 2: ", dia2)

	// Chamando a função do diaDaSemana para função 3 para fallthrough pois ele joga a clausula para a proxima verificação no switch
	dia3 := diaDaSemana3(1)
	fmt.Println("Dia da semana função 3: ", dia3)

}
