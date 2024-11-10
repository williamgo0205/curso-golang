package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

// Herança em Go vc declara o nome so struct mas não repassa o tipo dele (no cas o campk avaixo "pessoa")
type estudante struct {
	pessoa
	curso     string
	faculdade string
}

// 07 - Aula sobre Herança (Mas não exsite herança em GO e isso é o mais próximo) em go
func main() {
	fmt.Println("Herança")

	// Instanciando o struct "pessoa"
	pessoa := pessoa{"João", "Pedro", 20, 178}
	fmt.Println("pessoa: ", pessoa)

	// Instanciando o struct "estidante"
	estudante := estudante{
		pessoa:    pessoa,
		curso:     "Engenharia",
		faculdade: "USP"}
	fmt.Println("estudante: ", estudante)

	// Visualizando os campos do struct (não é preciso acessas esturdante.pessoa.nome e sim estudante.nome)
	fmt.Println("nome do estudante: ", estudante.nome)
	fmt.Println("sobrenome do estudante: ", estudante.sobrenome)
	fmt.Println("idade do estudante: ", estudante.idade)
	fmt.Println("altura do estudante: ", estudante.altura)
	fmt.Println("curso do estudante: ", estudante.curso)
	fmt.Println("faculdade do estudante: ", estudante.faculdade)
}
