package main

import "fmt"

// A partir do Go 1.18, é possível usar interfaces genéricas utilizando a palavra-chave "any" invés de "interface{}".

func generica(interf interface{}) {
	fmt.Println(interf)
}

func generica2(interf any) {
	fmt.Println(interf)
}

func main() {
	generica("String")
	generica(1)
	generica(true)

	fmt.Println("")

	generica2("String")
	generica2(1)
	generica2(true)

	fmt.Println("")

	fmt.Println(1, 2, "string", false, true, float64(12345))

	fmt.Println("")

	mapa := map[interface{}]interface{}{ // também pode ser map[any]any
		1:            "String",
		float32(100): true,
		"String":     "String",
		true:         float64(12),
	}

	fmt.Println(mapa)

}