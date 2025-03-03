package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

// 06 - Aula sobre structs em go
func main() {
	fmt.Println("Arquivo Structs")

	// Mock de endereço
	enderecoMock := endereco{"Rua Teste de exemplo", 97}

	// 1 - Declaracao e instanciacao do struct usuario
	var usuario1 usuario
	usuario1.nome = "William"
	usuario1.idade = 40

	fmt.Println("usuario 1: ", usuario1)

	// 2 - Declaracao e instanciacao do struct usuario por inferencia de tipos

	usuario2 := usuario{"Alessandra", 41, enderecoMock}
	fmt.Println("usuario 2: ", usuario2)

	// 3 - Declaracao e criação do usuario porme sem todos os dados
	usuario3 := usuario{nome: "Gabriel"}
	fmt.Println("usuario 3: ", usuario3)

	usuario4 := usuario{idade: 15}
	fmt.Println("usuario 4: ", usuario4)

	usuario5 := usuario{endereco: enderecoMock}
	fmt.Println("usuario 5: ", usuario5)
}
