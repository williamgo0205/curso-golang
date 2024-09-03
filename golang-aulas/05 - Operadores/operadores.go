package main

import "fmt"

// 05 - Aula sobre Operadores aritméticos em go
func main() {
	// Operadores Aritméticos
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 4 * 3
	mod := 10 % 2

	fmt.Println("")
	fmt.Println("*** Operadores Aritméticos ***")
	fmt.Println("soma: ", soma)
	fmt.Println("subtracao: ", subtracao)
	fmt.Println("divisao: ", divisao)
	fmt.Println("multiplicacao: ", multiplicacao)
	fmt.Println("resto da divisão: ", mod)

	// Operadores de Atribuição
	fmt.Println("")
	fmt.Println("*** Operadores de Atribuição ***")
	var variavel1 string = "string"
	variavel2 := "string2"

	fmt.Println("variavel1: ", variavel1)
	fmt.Println("variavel2: ", variavel2)

	// Operadores Relacionais
	fmt.Println("")
	fmt.Println("*** Operadores Relacionais ***")
	fmt.Println("1 > 2 verdadeiro ou falso? ", 1 > 2)
	fmt.Println("1 >= 2 verdadeiro ou falso? ", 1 >= 2)
	fmt.Println("1 == 2 verdadeiro ou falso? ", 1 == 2)
	fmt.Println("1 < 2 verdadeiro ou falso? ", 1 < 2)
	fmt.Println("1 <= 2 verdadeiro ou falso? ", 1 <= 2)
	fmt.Println("1 != 2 verdadeiro ou falso? ", 1 != 2)

	// Operadores Lógicos
	fmt.Println("")
	fmt.Println("*** Operadores Lógicos ***")
	// && (AND)
	fmt.Println("true && true verdadeiro ou falso? ", true && true)
	fmt.Println("true && false verdadeiro ou falso? ", true && false)
	fmt.Println("false && false verdadeiro ou falso? ", false && false)

	// || (OR)
	fmt.Println("true || true verdadeiro ou falso? ", true || true)
	fmt.Println("true || false verdadeiro ou falso? ", true || false)
	fmt.Println("false || false verdadeiro ou falso? ", false || false)

	// ! negacao
	fmt.Println("!true verdadeiro ou falso? ", !true)
	fmt.Println("!false verdadeiro ou falso? ", !false)

	// Operadores unários
	fmt.Println("")
	fmt.Println("*** Operadores Unários ***")

	numero := 10
	fmt.Println("número : ", numero)

	numero++ // O mesmo que numero +1
	fmt.Println("número++ : ", numero)

	numero += 15 // O mesmo que numero = numero + 15
	fmt.Println("numero += 15 :", numero)

	numero-- // O mesmo qe enumero -1
	fmt.Println("número-- : ", numero)

	numero -= 15 // O mesmo que numero = numero - 15
	fmt.Println("número -= 15 : ", numero)

	numero *= 3 // O mesmo que numero = numero * 3
	fmt.Println("número *= 3 : ", numero)

	numero /= 3 // O mesmo que numero = numero / 3
	fmt.Println("número /= 3 : ", numero)

	numero %= 3 // O mesmo que numero = numero % 3
	fmt.Println("número %= 3 : ", numero)

	// Operadores Ternários
	fmt.Println("")
	fmt.Println("*** Operadores Ternários ***")

	fmt.Println("Observação: Não possui operador ternário em GO precisa ser feito com IF e Else")

	numeroTeste := 7
	// texto := numeroTeste > 5 ? "numero maior que 5" : "numero menor que 5" (não existe em GO)

	var texto string
	if numeroTeste > 5 {
		texto = "numero maior que 5"
	} else {
		texto = "numero menor que 5"
	}

	fmt.Println(texto)

}
