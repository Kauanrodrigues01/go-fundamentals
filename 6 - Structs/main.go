package main

import "fmt"

// Definição da struct
type usuario struct {
	nome  string
	idade int
	email string
	cpf   string
}

// __str__ em Python, representação textual de um objeto em Golang
func (u usuario) String() string {
	return fmt.Sprintf("Nome: %s\nIdade: %d\nEmail: %s\nCPF: %s\n", u.nome, u.idade, u.email, u.cpf)
}

// Método com ponteiro para alterar o nome corretamente, sem criar uma cópia do objeto
// necessita de um ponteiro para alterar o valor do atributo, pois a função cria uma cópia do objeto e não altera o original, com o ponteiro eu referencio o objeto original e ele é alterado.
func (u *usuario) alterarNome(novoNome string) {
	u.nome = novoNome
}

func main() {
	exibirConteudo()

	usuario1 := usuario{nome: "João", idade: 30, email: "teste@example.com", cpf: "123.456.789-00"}
	fmt.Println(usuario1)

	usuario1.alterarNome("Maria") // Agora a alteração funciona corretamente!
	fmt.Println(usuario1)
}

var exibirConteudo = func() {
	fmt.Println("****** Structs em Golang ******")
	fmt.Println("")
}
