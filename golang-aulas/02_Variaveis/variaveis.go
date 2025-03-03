package main

import "fmt"

// Aula 02 - Declaração de tipos de variáveis em GO
func main() {
	// Declaração de variaveis

	// Variavel com tipagem explicita
	var variavel1 string = "Variável 1"
	fmt.Println(variavel1)

	// Variavel com tipagem implicita (oculta)
	variavel2 := "Variável 2"
	fmt.Println(variavel2)

	// Declaração de mais de um tipo de variável explicita
	var (
		variavel3 string = "Variável 3"
		variavel4 string = "Variável 4"
	)
	fmt.Println(variavel3, variavel4)

	// Declaração de mais de um tipo de variável inplicita
	variavel5, variavel6 := "Variável 5", "Variável 6"
	fmt.Println(variavel5, variavel6)

	// Declaração de constantes
	const constante1 string = "Constante 1"
	fmt.Println(constante1)

	// Declaração de mais de um tipo de comnstante
	const (
		constante2 string = "Constante 3"
		constante3 string = "Constante 4"
	)
	fmt.Println(constante2, constante3)

	// Exemplo Invertendo valores de variaveis
	var (
		variavel7 string = "Variável 7"
		variavel8 string = "Variável 8"
	)
	fmt.Println("valotres atuais: ", variavel7, variavel8)

	variavel7, variavel8 = variavel8, variavel7
	fmt.Println("valotres invertidos: ", variavel7, variavel8)

}
