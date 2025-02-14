package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Forma explicita de declarar variaveis
	var variavel1 string = "Variavel 1"
	// Forma implicita de declarar variaveis
	variavel2 := "Variavel 2"

	fmt.Println(variavel1, "-", reflect.TypeOf(variavel1))
	fmt.Println(variavel2, "-", reflect.TypeOf(variavel2))

	// Declarando variaveis em bloco
	var (
		variavel3 string = "Variavel 3"
		variavel4 string = "Variavel 4"
	)

	fmt.Println(variavel3, "-", reflect.TypeOf(variavel3))
	fmt.Println(variavel4, "-", reflect.TypeOf(variavel4))

	// Declarando variaveis em bloco de forma implicita
	variavel5, variavel6 := "Variavel 5", "Variavel 6"
	fmt.Println(variavel5, "-", reflect.TypeOf(variavel5))
	fmt.Println(variavel6, "-", reflect.TypeOf(variavel6))

	// Constantes
	const constente1 string = "Constante 1"
	fmt.Println(constente1, "-", reflect.TypeOf(constente1))

	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, "-", reflect.TypeOf(variavel5))
	fmt.Println(variavel6, "-", reflect.TypeOf(variavel6))

	// declarando variaveis sem atribuir valor
	var variavel7 string
	if variavel7 == "" {
		fmt.Println("Variavel7 não tem valor")
	}
}
