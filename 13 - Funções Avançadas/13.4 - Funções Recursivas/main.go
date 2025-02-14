package main

import "fmt"

func main() {
	fmt.Println("Funções Recursivas")
	fmt.Println(fibonacci(5))
}

func fibonacci(n uint) uint {
	if n <= 1 {
		return n
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}
