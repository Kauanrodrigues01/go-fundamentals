package main

import "fmt"

// O que são ponteiros?
// Ponteiros são variáveis que armazenam o endereço de memória de outra variável.
// Em Go, os ponteiros são representados por um asterisco (*) seguido do tipo da variável armazenada.
// Ponteiros são uteis quando queremos passar uma variável para uma função e alterar o valor dela.

// Quando usar ponteiros?
// 1. Para alterar o valor de uma variável dentro de uma função.
// 2. Para não copiar uma grande quantidade de dados ao passar uma variável para uma função.
// 3. Para passar um valor nulo para uma função.
// 4. Para compartilhar uma variável entre várias funções.

func main() {
	exibirConteudo()

	num1 := 10
	num2 := num1 // num2 é uma cópia de num1

	// Possuem endereco de memória diferentes
	fmt.Println("Endereço de memória de num1:", &num1)
	fmt.Println("Endereço de memória de num2:", &num2)
	fmt.Println("")


	// Ponteiro para num1
	num3 := &num1
	fmt.Println("Endereço de memória de num1:", &num1)
	fmt.Println("Endereço de memória de num3:", num3)
	fmt.Println("")

	// Como posso acessar o valor de num1 através do ponteiro num3?
	fmt.Println("Valor de num1:", num1)
	fmt.Println("Valor de num1 através do ponteiro num3:", *num3)
	fmt.Println("")

	// Se eu alterar o valor de num1 através do ponteiro num3, o valor de num1 também será alterado? Sim!
	*num3 = 20
	fmt.Println("Valor de num1:", num1)
	fmt.Println("Valor de num1 através do ponteiro num3:", *num3)
	fmt.Println("")

	// Utilizando nil para ponteiros
	var ponteiroNulo *int
	fmt.Println("Ponteiro nulo:", ponteiroNulo)
	fmt.Println("")
}

var exibirConteudo = func() {
	println("****** Ponteiros em Golang ******")
	println("")
}