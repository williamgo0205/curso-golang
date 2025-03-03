package main

import "fmt"

// 12 - Aula sobre Estruturas de controle em GO
func main() {
	fmt.Println("Estruturas de Controle")

	numero := 10

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	numero = -20

	// Declarando uma variável dentro do if (chamada - if init)
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Maior que 0")
	} else if outroNumero < -10 {
		fmt.Println("Menor que -10")
	} else {
		fmt.Println("Número entre 0 e -10")
	}
}
