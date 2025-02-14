package main

import "fmt"

func main() {
	exibirConteudo()

	numero := -5

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	// declaração de variável dentro do if, essa variável só existe dentro do bloco do if, não é acessível fora dele
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Número é maior que zero")
	} else if numero < -10 {
		fmt.Println("Número é menor que -10")
	} else {
		fmt.Println("Entre 0 e -10")
	}

}

var exibirConteudo = func() {
	fmt.Println("****** If e Else em Golang ******")
	fmt.Println("")
}
