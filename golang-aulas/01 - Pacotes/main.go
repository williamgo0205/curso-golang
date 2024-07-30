// Para considerar um arquivo executável e necessário na primeira linha escrever "package main"
package main

import (
	"fmt"
	"modulo/auxiliar"
)

// Para ser um arquivo executavel alem do "package main" também é necessario existir a funcao "main"
// como abaixo
func main() {
	fmt.Println("Escrevendo do arquivo main")
	auxiliar.Escrever()
}
