package main

import (
	"fmt"
)

// Aula 04 - Funcoes em GO

// Função com retorno
func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

// Função com mais de um retorno
// Exemplo: func <nome da função> (parâmeros de entrada) (retornos da função)
func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	// 1 - Chamando a funcao somar repassando dois valores
	soma := somar(10, 20)
	fmt.Println(soma)

	// 2 - Declarando uma função para uma variável e chamando essa funçao sequencialmente
	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	fmt.Println(f("Texto da Função 1"))

	// 3 - Chamando uma função com mais de um retorno
	resultadoSoma, resultadoSubtracao := calculosMatematicos(20, 10)
	fmt.Println("resultado do retorno da função calculosMatematicos para dois valores de retorno: soma = ", resultadoSoma, " subtracao = ", resultadoSubtracao)

	// 3.1 - Chamando uma função com mais de um retorno porém sem utilizar um dos valores, para isso utiliza-se o "_"
	resultadoSoma2, _ := calculosMatematicos(33, 15)
	fmt.Println("resultado do retorno da função calculosMatematicos somente para um valor de retorno: soma = ", resultadoSoma2)
}
