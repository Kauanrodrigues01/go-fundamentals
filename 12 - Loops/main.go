package main

import (
	"fmt"
	"time"
)

// Loops em Go
// Go possui apenas um tipo de loop, o for.
// O for pode ser usado de quatro formas:
// 1. Com uma única condição;
// 2. Com uma inicialização, condição e pós;
// 3. Com a palavra chave range.
// 4. Loop infinito.


func main() {
	// 1. Com uma única condição
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	fmt.Println("-----------")

	// 2. Com uma inicialização, condição e pós
	for i := 10; i < 20; i++ {
		fmt.Println(i)
	}

	fmt.Println("-----------")

	// 3. Com a palavra chave range
	// A palavra chave range é usada para iterar sobre um slice, array, string, map ou channel.
	for i := range 10 {
		fmt.Println(i)
	}

	fmt.Println("-----------")

	for index, value := range "Hello" {
		fmt.Println(index, "-", string(value))
	}

	fmt.Println("-----------")

	for index, value := range []int{1, 2, 3, 4, 5} {
		fmt.Println(index, "-", value)
	}

	fmt.Println("-----------")

	for key, value := range map[string]string{"a": "apple", "b": "banana"} {
		fmt.Println(key, "-", value)
	}

	fmt.Println("-----------")

	// 4. Loop infinito
	// Para criar um loop infinito, basta não passar nenhuma condição. E para sair do loop, basta usar a palavra chave break.
	i = 0
	for {
		fmt.Println("Loop infinito")
		fmt.Println(time.Now())
		i++
		if i == 20 {
			break
		}
	}

}