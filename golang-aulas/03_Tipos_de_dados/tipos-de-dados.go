package main

import (
	"errors"
	"fmt"
)

// Aula 03 - Tipos de Dados em GO
func main() {

	// ** Numeros Inteiros ***
	// Tipos:
	// int (usa a arquitetura do computador),
	// int8 (8 bits),
	// int16 (16 bits),
	// int32 (32 bits),
	// int64 (64 bits)
	var numero int = -1000000000000000000
	fmt.Println(numero)

	// uint (Unsygned Int) - Int sem sinal
	// uint (usa a arquitetura do computador),
	// uint8 (8 bits),
	// uint16 (16 bits),
	// uint32 (32 bits),
	// uint64 (64 bits)
	var numero2 int = 10000
	fmt.Println(numero2)

	// Alias para o int32 é rune para números que representa caracteres da tabela ASCII
	//var numero3 int32 = 123456
	var numero3 rune = 123456
	fmt.Println(numero3)

	// Alias para o uint8 é byte
	//var numero3 8 = 123
	var numero4 byte = 123
	fmt.Println(numero4)

	// *** Numeros Reais ***
	// Tipos:
	// float32 - Numeros reais de 32 bits
	// float64 - Numeros reais para 64 bits
	var numeroReal float32 = 1.23456
	fmt.Println(numeroReal)

	var numeroReal2 float64 = 12355668.23456252
	fmt.Println(numeroReal2)

	numeroReal3 := 1235.23
	fmt.Println(numeroReal3)

	// *** Strings ***
	// Tipos:
	// Declaração sempreem aspas duplas
	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	// Nao existe o tipo CHAR para caractere em Go
	// Quando colocado em Aspas simples é considerado um INT no caso da tabela ASCII
	char := 'W'
	fmt.Println(char)

	// *** Booleano ***
	// bool - true/false
	var booleano1 bool = true
	fmt.Println(booleano1)

	var booleano2 bool = false
	fmt.Println(booleano2)

	// *** Erro ***
	// Tipo dele é ERROR
	var erro error = errors.New("Erro de execução do teste")
	fmt.Println(erro)

	// *** Valor zero ***
	// Será atribuido o valor zero a uma variável quando, essa variável não possuir valor inicial
	// No caso string será um valor vazio
	// No caso de numero será "0"
	// No caso do booleano é FALSE
	// No caso do error é <nil> (nulo)
	var valorZero string
	fmt.Println("valor zero string:", valorZero)

	var valorZero1 int16
	fmt.Println("valor zero int:", valorZero1)

	var valorZero2 bool
	fmt.Println("valor zero booleano:", valorZero2)

	var valorZero3 error
	fmt.Println("valor zero error:", valorZero3)
}
