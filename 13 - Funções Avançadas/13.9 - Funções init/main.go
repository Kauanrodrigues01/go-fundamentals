package main

import "fmt"

// init é uma função que é executada antes da função main
// é muito utilizada para inicializar variáveis globais
// é executada apenas uma vez
// não aceita parâmetros e não retorna valores

func main() {
	fmt.Println("Função main")
}

func init() {
	fmt.Println("Função init")
}