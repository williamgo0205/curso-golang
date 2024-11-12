package main

import (
	"fmt"
	"reflect"
)

// 09- Aula sobre Arrays e Slices em go
func main() {
	fmt.Println("Arrays e Slices")

	// ARRAY - Array é uma lista de valores

	// Exemplo 1
	var array1 [5]string
	array1[0] = "Posição 1"
	array1[1] = "Posição 2"
	array1[2] = "Posição 3"
	array1[3] = "Posição 4"
	array1[4] = "Posição 5"

	fmt.Println("Array1: ", array1)

	// Exemplo 2
	array2 := [5]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5"}
	fmt.Println("Array2: ", array2)

	// Exemplo 3 - Fixa o tamanho do array de acrdo com a quantodade de posições informadas no array, nesse caso um array de 9 posições
	array3 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Array3: ", array3)

	// SLICES - O Slice não é um array ele não possui o tamanho fixo como no array, podendo mudar

	// Exemplo 1
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("slice: ", slice)

	// Adicionando um novo item nos lice
	slice = append(slice, 6)
	fmt.Println("slice adicionado: ", slice)

	// Exermplo 2
	// É possível obter apenas uma parte de um array no Slice
	// Uma ressalva quanto a monenclarura: "array3[0:3]",significa que irá buscar o índice "0" (inclusivo), até o índice "3" (exclusivo),
	// porém, o indice 3 nãos erá carregado pois ele é exclusivo, ou seja ele exclui esse indice, esendo assim exibira [1 2 3] e não [1 2 3 4]
	slice2 := array3[0:3]
	fmt.Println("slice2: ", slice2)

	array3[1] = 27
	fmt.Println("slice2 após alteração no array3: ", slice2)

	// Verificando os tipos de slice e array para demonstrar que sao diferentes
	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

}
