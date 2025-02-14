package main

import "fmt"

// Defer = adiar, postergar, atrasar
// Defer é chamado por ultimo dentro de um bloco de código
// Executa uma função no momento que a função atual está sendo finalizada
// A função defer é executada por último
// É muito utilizado para fechar conexões com banco de dados, fechar arquivos, etc
// Defer é executado sempre, independente de ter um return antes
// Defer é uma pilha, ou seja, o último defer a ser chamado será o primeiro a ser executado
// Defer é muito utilizado para tratamento de erros, pois ele é executado sempre, mesmo que ocorra um erro
// Defer é muito utilizado para tratamento de logs, pois ele é executado sempre, mesmo que ocorra um erro

func funcao1() {
	fmt.Println("Executando a função 1")
}

func funcao2() {
	fmt.Println("Executando a função 2")
}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer fmt.Println("Média calculada. Resultado será retornado!")
	fmt.Println("Entrando na função para verificar se o aluno está aprovado")

	media := (n1 + n2) / 2
	defer fmt.Println("A média do aluno é:", media)

	if media >= 6 {
		return true
	}
	return false
}

func main() {
	fmt.Println(alunoEstaAprovado(7, 8))
}