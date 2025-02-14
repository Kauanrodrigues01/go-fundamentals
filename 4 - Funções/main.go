package main

import "fmt"

func main() {
	exibirConteudo()

	soma := somar(8, 9)
	fmt.Println(soma)

	subtracao := subtrair(8, 9)
	fmt.Println(subtracao)


	resultadoSoma, resultadoSubtracao := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma, resultadoSubtracao)

	resultadoSubtracao2 := calcular(subtrair, 8, 9)
	fmt.Println(resultadoSubtracao2)

}

// Funções simples com retorno
func subtrair(n1 int8, n2 int8) int8 {
	return n1 - n2
}

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

// Função com retorno multiplo
func calculosMatematicos(n1 int8, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

// Função como parâmetro
func calcular(f func(int8, int8) int8, n1 int8, n2 int8) int8 {
	return f(n1, n2)
}

// Função ANÔNIMA
var exibirConteudo = func() {
	fmt.Println("******** Funções em Golang ********")
	fmt.Println("1 - Definição de Funções")
	fmt.Println("2 - Funções com retorno simples")
	fmt.Println("3 - Funções com retorno múltiplo")
	fmt.Println("4 - Funções como parâmetro")
	fmt.Println("************************************")
	fmt.Println("")
}