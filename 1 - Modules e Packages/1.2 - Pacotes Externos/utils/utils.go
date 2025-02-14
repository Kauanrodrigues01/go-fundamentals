package utils

import (
    "fmt"
    "strconv"
    "time"
)

func PegarSaudacao() string {
    var saudacao string

    horaAtualStr := time.Now().Format("15")

    horaAtual, err := strconv.Atoi(horaAtualStr)
    if err != nil {
        fmt.Println("Erro ao converter a hora:", err)
        return "Olá" // Saudação padrão em caso de erro
    }

    if horaAtual >= 6 && horaAtual < 12 {
        saudacao = "Bom dia"
    } else if horaAtual >= 12 && horaAtual < 18 {
        saudacao = "Boa tarde"
    } else {
        saudacao = "Boa noite"
    }

    return saudacao
}
