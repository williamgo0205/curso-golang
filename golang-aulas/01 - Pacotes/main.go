// Para considerar um arquivo executável e necessário na primeira linha escrever "package main"
package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

// Para ser um arquivo executavel alem do "package main" também é necessario existir a funcao "main"
// como abaixo
func main() {
	fmt.Println("Escrevendo do arquivo main")
	auxiliar.Escrever()

	// Funcao para testar o pacote "github.com/badoux/checkmail"
	erro := checkmail.ValidateFormat("devbook@gmail.com")
	fmt.Println(erro)
}
