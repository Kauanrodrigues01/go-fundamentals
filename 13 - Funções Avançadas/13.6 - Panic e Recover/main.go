package main

import "fmt"

// PANIC
// 1. panic é uma função que interrompe a execução do programa e exibe uma mensagem de erro
// 2. Quando o panic é chamado, ele para a execução do programa e exibe uma mensagem de erro, apenas as funções defer serão executadas antes do programa ser encerrado
// 3. O panic é uma forma de parar a execução do programa quando algo inesperado acontece, como uma divisão por zero ou um acesso a um índice inválido de um slice

// RECOVER
// 1. recover é uma função que retoma o controle da execução do programa após um panic
// 2. O recover só funciona dentro de uma função defer
// 3. O recover só consegue capturar um panic que acontece na mesma goroutine em que ele está sendo executado

// EXEMPLO
// 1. Vamos criar uma função que verifica se um aluno está aprovado ou não
// 2. Se a média do aluno for maior que 6, ele está aprovado
// 3. Se a média do aluno for menor que 6, ele está reprovado
// 4. Se a média do aluno for exatamente 6, vamos lançar um panic
// 5. Vamos criar uma função main que chama a função alunoEstaAprovado e vamos tratar o panic com recover

// OBSERVAÇÕES:
// 1. O defer só funciona se ele for declarado antes do panic


func main() {
	aprovado := alunoEstaAprovado(6, 6) // panic
	println(aprovado)

	fmt.Println("Testando se a função main continua sendo executada, mesmo após um panic, tratando o erro com recover")
}

func recuperExecucao() {
	if r := recover(); r != nil { // r é a mensagem de erro que foi passada para o panic, se não for nil, quer dizer que um panic foi lançado
		fmt.Println("Ocorreu um erro:", r)
	}
}

func alunoEstaAprovado(n1, n2 float64) bool {
	defer recuperExecucao()

	media := (n1 + n2) / 2
	if media > 6 {
		return true
	}
	if media < 6 {
		return false
	}

	panic("A MÉDIA É EXATAMENTE 6!")
}