package main

import "fmt"

func main() {
	testandoFuncaoSemPonteiro()

	testandoFuncaoComPonteiro()
}

// Essa função recebe apenas uma copia do valor da variável numero, mas não o endereço de memória, ou seja, não é possível alterar o valor da variável numero, apenas a copia que foi passada para a função
func inverterSinal(numero int) int {
	return numero * -1
}

func testandoFuncaoSemPonteiro() {
	fmt.Println("Testando função sem ponteiro:")
	numero := 20
	numeroInvertido := inverterSinal(numero)
	fmt.Println("Número invertido (copia):", numeroInvertido)
	fmt.Println("Número original:", numero)
	fmt.Println("")
	// Observamos que o valor da variável numero não foi alterado, pois a função inverterSinal recebeu apenas uma copia do valor da variável numero
}


func inverterSinalComPonteiro(numero *int) int {
	*numero = *numero * -1 // Altera o valor da variável numero, no endereço de memória que foi passado para a função
	return *numero
}

func testandoFuncaoComPonteiro() {
	fmt.Println("Testando função com ponteiro")
	numero := 20
	numeroInvertido := inverterSinalComPonteiro(&numero)
	fmt.Println("Número invertido (copia):", numeroInvertido)
	fmt.Println("Número original:", numero)
	// Observamos que o valor da variável numero foi alterado, pois a função inverterSinalComPonteiro recebeu o endereço de memória da variável numero
}