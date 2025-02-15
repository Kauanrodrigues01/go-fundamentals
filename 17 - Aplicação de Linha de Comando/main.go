package main

import (
	"fmt"
	"linha-de-comando/app"
	"log"
	"os"
)

func main() {
	fmt.Println("Iniciando a aplicação...")

	app := app.Gerar()
	erro := app.Run(os.Args)
	if erro != nil {
		log.Fatal(erro)
	}
}