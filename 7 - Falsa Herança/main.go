package main

import "fmt"

type pessoa struct {
	nome  string
	idade uint
}

func (p pessoa) saudacao() {
	fmt.Println("Olá, meu nome é", p.nome)
}

// Isso é o mais próximo que Go tem de herança, que é a composição de structs, mas não é herança de fato
type estudante struct {
	pessoa    // embutindo a struct pessoa
	curso     string
	faculdade string
}

type professor struct {
	pessoa         // embutindo a struct pessoa
	especializacao string
	salario        float64
}

func main() {
	exibirConteudo()

	// Inicializando a struct estudante
	estudante1 := estudante{
		pessoa: pessoa{
			nome:  "João",
			idade: 20,
		},
		curso:     "Engenharia",
		faculdade: "USP",
	}

	// Imprimindo o estudante
	fmt.Println(estudante1.nome)
	fmt.Println(estudante1.idade)
	fmt.Println(estudante1.curso)
	fmt.Println(estudante1.faculdade)
	estudante1.saudacao()
	fmt.Println("")

	professor1 := professor{
		pessoa: pessoa{
			nome: "Maria",
			idade: 30,
		},
		especializacao: "Engenharia de Software",
		salario: 5000.00,
	}

	// Imprimindo o professor
	fmt.Println(professor1.nome)
	fmt.Println(professor1.idade)
	fmt.Println(professor1.especializacao)
	fmt.Println(professor1.salario)
	professor1.saudacao()
}

var exibirConteudo = func ()  {
	fmt.Println("****** Falsa Herança em Golang ******")
	fmt.Println("")
}