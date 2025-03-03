package main

import "fmt"

// 11 - Aula sobre Map
func main() {
	fmt.Println("Aula de MAP em GO")

	// Declaração: map[{tipo da chave}]{tipo dos valores}
	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}

	fmt.Println("usuário: ", usuario)
	fmt.Println("usuário verificando o campo nome: ", usuario["nome"])
	fmt.Println("usuário verificando o campo sobrenome: ", usuario["sobrenome"])

	// Declaração MAP anilnhado
	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "João",
			"ulltimo":  "Pedro",
		},
		"curso": {
			"curso":  "Engenharia",
			"campus": "Camnpus 1",
		},
	}

	fmt.Println("usuário 2: ", usuario2)
	fmt.Println("usuário verificando o campo nome: ", usuario2["nome"])
	fmt.Println("usuário verificando o campo curso: ", usuario2["curso"])

	// Deletar uma chave do map
	// Estrutura delete({nome do map}, {nome da chave})
	delete(usuario2, "nome")

	fmt.Println("usuário 2 deletado: ", usuario2)

	// Adicionar uma chave do map
	// Estrutura deve seguir a estrutura original do map, neste caso map[string]string
	// {nome do map}["{nova chave}"]{tipo}
	usuario2["signo"] = map[string]string{
		"nome": "Touro",
	}

	// Verificando valores do campo noco adicionado.
	fmt.Println("usuário 2 adicionado: ", usuario2)
	fmt.Println("usuário verificando o campo signo: ", usuario2["signo"])

}
