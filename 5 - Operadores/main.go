package main

import "fmt"

func main() {
	exibirConteudo()
	// Operadores Aritméticos (+, -, *, /, %)
	soma := 10 + 5
	subtracao := 10 - 5
	multiplicacao := 10 * 5
	divisao := 10 / 5
	restoDaDivisao := 10 % 5
	fmt.Println(soma, subtracao, multiplicacao, divisao, restoDaDivisao)
	fmt.Println("")

	// Operadores de Atribuição (=, +=, -=, *=, /=, %=)
	var variavel1 int = 10
	variavel1 += 5 // variavel1 = variavel1 *operador* 5
	variavel1 -= 5
	variavel1 *= 5
	variavel1 /= 5
	variavel1 %= 5
	fmt.Println(variavel1)
	fmt.Println("")

	// operadores de comparação/relacionais (==, !=, >, <, >=, <=)
	fmt.Println(10 == 10)
	fmt.Println(10 != 10)
	fmt.Println(10 > 10)
	fmt.Println(10 < 10)
	fmt.Println(10 >= 10)
	fmt.Println(10 <= 10)
	fmt.Println("")

	// Operadores Lógicos (&&, ||, !)
	fmt.Println(true && true) // and
	fmt.Println(true || true) // or
	fmt.Println(!true) // not
}

var exibirConteudo = func() {
	fmt.Println("****** Operadores em Golang ******")
	fmt.Println("1. Operadores Aritméticos")
	fmt.Println("2. Operadores de Atribuição")
	fmt.Println("3. Operadores de Comparação/Relacionais")
	fmt.Println("4. Operadores Lógicos")
	fmt.Println("***********************************")
	fmt.Println("")
}