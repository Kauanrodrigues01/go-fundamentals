package main

import (
	"fmt"
	"math"
)

// A interface "forma" define um comportamento que qualquer estrutura pode implementar.
// Para ser considerada do tipo "forma", a estrutura deve ter um método "area()" que retorna um float64.
type forma interface {
	area() float64
}

// Definição da estrutura "circulo" com um campo "raio".
type circulo struct {
	raio float64
}

// Implementação do método "area()" para "circulo".
// Calcula a área usando a fórmula: π * raio².
func (c circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

// Definição da estrutura "retangulo" com os campos "altura" e "largura".
type retangulo struct {
	altura  float64
	largura float64
}

// Implementação do método "area()" para "retangulo".
// Calcula a área usando a fórmula: altura * largura.
func (r retangulo) area() float64 {
	return r.altura * r.largura
}

// Função que recebe qualquer "forma" e imprime sua área.
// Como "circulo" e "retangulo" possuem o método "area()", eles podem ser usados aqui.
func escreverArea(f forma) {
	area := f.area()
	formatted := fmt.Sprintf("%.2f", area) // Formata para duas casas decimais.
	fmt.Println("A área da forma é:", formatted)
}

func main() {
	c := circulo{raio: 10}       // Criando um círculo com raio 10.
	r := retangulo{altura: 10, largura: 15} // Criando um retângulo de 10x15.

	escreverArea(c) // Chamando a função com um círculo, que por sua vez chama o método "area()" de "circulo".
	escreverArea(r) // Chamando a função com um retângulo, que por sua vez chama o método "area()" de "retangulo".
}

// Para chamar os metodos escreverArea com um retangulo ou circulo e isso funcionar, é necessário que os structs de circulo e retangulo implementem o método area() que retorna um float64, pois a interface forma exige isso e a função escreverArea chama esse método "area()".