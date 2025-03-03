package auxiliar

import "fmt"

// Necessario escrever o nome dasfuncoes com a primeira letra maiuscula pois o GO
// nao é uma linguagem orientada a objeto e por tal motivo, nao reconhece metodos publicos/privados/protected
// sendo assim, para uma funcao ser visível fora do pacote que ela se encontra, somente se ela for escrita
// com a primeira letra maiuscula, ela sendo toda minuscula, somente será visivel dentro do pacote atual
func escrever2() {
	fmt.Println("Escrevendo do pacote auxiliar pela função 2")
}
