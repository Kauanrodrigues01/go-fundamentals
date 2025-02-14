package main

import "fmt"

// Qual a diferença entre arrays e slices?
// Arrays são estruturas de dados de tamanho fixo, enquanto slices são estruturas de dados de tamanho variável.
// Arrays são mais rápidos e eficientes, enquanto slices são mais flexíveis e fáceis de usar.

func main() {
	exibirConteudo()

	// ARRAYS
	var array1 [5]int // Array de 5 posições do tipo int, tem tamanho fixo
	array1[0] = 10
	array1[1] = 20
	fmt.Println("Array1:", array1)
	fmt.Println("")

	var array2 = [5]string{"a", "b", "c", "d", "e"} // Array de 5 posições do tipo string, tem tamanho fixo
	fmt.Println("Array2:", array2)
	fmt.Println("")

	var array3 = [...]int{1, 2, 3, 4, 5} // O compilador conta o tamanho do array e atribui automaticamente, ele analisa o número de elementos que foram inicializados e define o tamanho do array com base nisso
	fmt.Println("Array3:", array3)
	fmt.Println("")


	// SLICES
	var slice1 = []int{10, 20, 30, 40, 50} // Slice do tipo int, tem tamanho variável
	slice1 = append(slice1, 60) // Adicionando um novo elemento ao slice
	fmt.Println("Slice1:", slice1)
	fmt.Println("")


	// Criando um slice a partir de um array
	var slice2 = array1[0:3] // Cria um slice com os elementos de array1 nas posições 0, 1 e 2
	slice2 = append(slice2, 70) // Adicionando um novo elemento ao slice
	fmt.Println("Slice2:", slice2)
	fmt.Println("")


	// ARRAYS INTERNOS
	slice3 := make([]int, 5, 10) // Cria um slice com 5 elementos por default e capacidade para ter até 10 elementos, funciona como um limitador para o slice
	fmt.Println("Slice3 (Antes):", slice3)
	slice3 = append(slice3, 10, 20, 30, 40, 50) // Adicionando 5 elementos ao slice
	fmt.Println("Slice3 (Depois):", slice3)
	fmt.Println("")

	// Como descobir o tamanho e a capacidade de um slice?
	slice4 := make([]int, 5, 10)
	fmt.Println("Tamanho do slice4:", len(slice4))
	fmt.Println("Capacidade do slice4:", cap(slice4))
	fmt.Println("")

	// ultrapassando a capacidade do slice
	slice4 = append(slice4, 60, 70, 80, 90, 100, 110, 120, 130, 140, 150)
	fmt.Println("Tamanho do slice4 após ultrapassar a capacidade:", len(slice4))
	fmt.Println("Capacidade do slice4 após ultrapassar a capacidade:", cap(slice4))
}

var exibirConteudo = func() {
	fmt.Println("****** Arrays e slices em Golang ******")
	fmt.Println("")
}