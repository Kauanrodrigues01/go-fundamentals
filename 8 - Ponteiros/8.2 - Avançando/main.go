package main

import "fmt"

// Definindo uma struct chamada 'Pessoa'
type Pessoa struct {
	nome  string
	idade int
}

// Função que altera a idade de uma Pessoa através de um ponteiro
func alterarIdade(p *Pessoa) {
	p.idade = 40
}

func main() {
	// Criando uma instância de Pessoa
	pessoa1 := Pessoa{nome: "João", idade: 30}

	// Mostrando o conteúdo da pessoa1
	fmt.Println("Pessoa1:", pessoa1)

	// Criando um ponteiro para pessoa1
	ponteiroPessoa := &pessoa1

	// Alterando o valor de 'nome' através do ponteiro
	ponteiroPessoa.nome = "Maria"

	// Exibindo o conteúdo da struct após a alteração
	fmt.Println("Pessoa1 após alteração:", pessoa1)

	// Passando a struct para uma função
	alterarIdade(ponteiroPessoa)

	// Exibindo o conteúdo da struct após alteração na função
	fmt.Println("Pessoa1 após alteração de idade:", pessoa1)
}
