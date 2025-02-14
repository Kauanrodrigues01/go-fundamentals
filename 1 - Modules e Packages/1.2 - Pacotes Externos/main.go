package main

import (
	"fmt"
	"1-pacote/utils"
	"github.com/badoux/checkmail"
	"github.com/fatih/color"
)

func main() {
	saudacao := utils.PegarSaudacao()

	var email string
	fmt.Print(saudacao + ". Seja bem-vindo ao programa de verificação de email.\n")
	fmt.Println("Digite seu email:")
	fmt.Scanln(&email)

	err := checkmail.ValidateFormat(email)

	if err != nil {
		color.Red("Formato de email inválido")
	} else {
		color.Green("Tudo ok!")
	}
}