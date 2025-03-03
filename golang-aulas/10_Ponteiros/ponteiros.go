package main

import "fmt"

// 10 - Aula sobre Ponteiros em GO
func main() {
	// 1 - Atribuiçao de valores por cópia
	// Declaração de variáveis
	var variavel1 int = 10
	var variavel2 int = variavel1

	fmt.Println("variavel 1: ", variavel1)
	fmt.Println("variavel 2: ", variavel2)

	// Adicionando mais um valor a variável 1
	variavel1++
	fmt.Println("variavel 1: ", variavel1)

	// 2 - Atribuição de valores utilizando ponteiros
	// Ponteiro pe uma referencia de memória

	var variavel3 int
	var ponteiro *int

	fmt.Println("Valores sem atribuição")
	fmt.Println("variavel 3: ", variavel3)
	fmt.Println("Ponteiro: ", ponteiro)

	variavel3 = 100
	ponteiro = &variavel3

	fmt.Println("Atribuindo valores iniciais")
	fmt.Println("variavel 3: ", variavel3)
	fmt.Println("Ponteiro endereço de memória: ", ponteiro)
	fmt.Println("Ponteiro desreferenciação: ", *ponteiro)

	variavel3 = 150

	fmt.Println("Verificando valores alterados")
	fmt.Println("variavel 3: ", variavel3)
	fmt.Println("Ponteiro desreferenciação: ", *ponteiro)
}
