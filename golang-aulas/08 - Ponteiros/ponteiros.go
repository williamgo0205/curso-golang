package main

import "fmt"

// 08 - Aula sobre Ponteiros em go
func main() {
	fmt.Println("ponteiros")

	// 1 - Atribuição via valores
	var variavel1 int = 10
	var variavel2 int = variavel1

	fmt.Println("Teste 1")
	fmt.Println("variavel1: ", variavel1, ", variavel2: ", variavel2)

	variavel1++
	fmt.Println("Teste 2")
	fmt.Println("variavel1: ", variavel1, ", variavel2: ", variavel2)

	// 2 - Atribuição via ponteiro - Ponteiro é Referência de memória
	var variavel3 int
	var ponteiro *int

	fmt.Println("Teste 3")
	fmt.Println("variavel3: ", variavel3, ", ponteiro: ", ponteiro)

	variavel3 = 100
	// Adicionando o "&" você armazena o endereçode memória da variável desejada, não seu valor mas seu endereço de memória
	ponteiro = &variavel3

	fmt.Println("Teste 4")
	fmt.Println("variavel3: ", variavel3, ", ponteiro - endereço de memoria: ", ponteiro)

	// Adicionando um "*" na frente do ponteiro você obtem o valor dele e nao o endereço de memoria (desreferenciação)
	fmt.Println("variavel3: ", variavel3, ", ponteiro - valor: ", *ponteiro)

	variavel3++
	fmt.Println("Teste 5")
	fmt.Println("variavel3: ", variavel3, ", ponteiro - endereço de memoria: ", ponteiro)
	fmt.Println("variavel3: ", variavel3, ", ponteiro - valor: ", *ponteiro)
}
