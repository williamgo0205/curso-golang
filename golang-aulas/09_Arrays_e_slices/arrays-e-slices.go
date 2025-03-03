package main

import (
	"fmt"
	"reflect"
)

// 09- Aula sobre Arrays e Slices em go
func main() {
	fmt.Println("---------------")
	fmt.Println("Arrays e Slices")
	fmt.Println("---------------")
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

	fmt.Println("---------------")
	fmt.Println("Arrays Internos")
	fmt.Println("---------------")

	// Arrays Internos

	// make é uma função que ocupa um espaço na memória para uma determinada necessidade
	// Essa função recebe 3 parâmetros:
	//    1 - Tipo do objeto (int, string...)
	//    2 - Tamanho do objeto
	//    3 - capacidade (quantidade máxima dos itens)
	slice3 := make([]float32, 10, 11)
	fmt.Println("slice3: ", slice3)
	fmt.Println("slice3 tamanho (len): ", len(slice3))
	fmt.Println("slice3 capacidade (cap): ", cap(slice3))

	// Adicionando um valor no slice
	slice3 = append(slice3, 7)
	fmt.Println("slice3: ", slice3)
	fmt.Println("slice3 tamanho (len): ", len(slice3))
	fmt.Println("slice3 capacidade (cap): ", cap(slice3))

	// Adicionando mais um valor no slice
	slice3 = append(slice3, 9)
	fmt.Println("slice3: ", slice3)
	fmt.Println("slice3 tamanho (len): ", len(slice3))
	fmt.Println("slice3 capacidade (cap): ", cap(slice3))

	// Adicionando um slace através da opção make sem o valor de capacidade qe é opcional
	slice4 := make([]float32, 5)
	fmt.Println("slice4: ", slice4)
	fmt.Println("slice4 tamanho (len): ", len(slice4))
	fmt.Println("slice4 capacidade (cap): ", cap(slice4))

	// Adicionando um valor no slice
	slice4 = append(slice4, 10)
	fmt.Println("slice4: ", slice4)
	fmt.Println("slice4 tamanho (len): ", len(slice4))
	fmt.Println("slice4 capacidade (cap): ", cap(slice4))

	// Resumidamente:
	// Array - é uma lista com o tamanho fixo
	// Slice - é uma lista sem tamanho fixo
}
