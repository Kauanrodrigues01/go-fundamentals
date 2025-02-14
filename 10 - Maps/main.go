package main

import "fmt"

// O que são maps?
// Maps são estruturas de dados que armazenam pares chave-valor, onde cada chave é única e não pode ser repetida.
// São como dicionários em Python.

func main() {
	exibirConteudo()

	// MAP convencional
	usuario := map[string]string{
		"nome":      "João",
		"sobrenome": "da Silva",
	}
	fmt.Println(usuario)
	fmt.Println(usuario["nome"])

	// MAP com valor MAP
	usuario2 := map[string]map[string]string{
		"aluno": {
			"nome":      "Maria",
			"sobrenome": "da Silva",
		},
		"curso": {
			"nome": "Engenharia de Software",
			"carga_horaria": "3000h",
			"professor": "João da Silva",
		},
	}

	fmt.Println("Usuário 2(antes de deletar um item):", usuario2)
	fmt.Println("")

	delete(usuario2, "curso") // Deletando um item do map
	fmt.Println("Usuário 2(depois de deletar um item):", usuario2)
	fmt.Println("")

	// Adicionando um novo item ao map
	usuario2["status"] = map[string]string{
		"ativo": "true",
		"matriculado": "true",
		"aprovado": "false",
	}
	fmt.Println("Usuário 2(depois de adicionar um item):", usuario2)
	fmt.Println("")


	// MAP com valor SLICE
	usuario3 := map[string][]string{
		"aluno": {"Maria", "da Silva"},
		"curso": {"Engenharia de Software", "3000h", "João da Silva"},
	}
	fmt.Println("Usuário 3:", usuario3)
	fmt.Println("")

	// MAP com valor de chave diferente do valor
	usuario4 := map[string]int{
		"idade": 30,
		"peso":  70,
	}
	fmt.Println("Usuário 4:", usuario4)
}

var exibirConteudo = func() {
	fmt.Println("****** Maps em Golang ******")
	fmt.Println("")
}
