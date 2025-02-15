package utils

import (
	"strings"
)

func FormatarEnderecoHost(host *string) {
	if strings.Index(*host, "http://") == 0 {
		*host = strings.TrimPrefix(*host, "http://")
	} else if strings.Index(*host, "https://") == 0 {
		*host = strings.TrimPrefix(*host, "https://")
	} else if strings.Index(*host, "www.") == 0 {
		*host = strings.TrimPrefix(*host, "www.")
	}

	// strings são um conjunto de bytes, não de caracteres, então quando pega um valor usando um indice, o go retorna um byte e não um caractere, para transformar um byte/rune em string é necessário usar o "string(byte/rune)""

	if len(*host) > 0 && (*host)[len(*host)-1] == '/' { // '/' é um byte ou rune e não uma string
		*host = (*host)[:len(*host)-1]
	}

	// Se usar "*host[len(*host)-1]" dá erro porque o go primeiro tenta acessar o valor do indice e só depois tenta desreferenciar o ponteiro

	// utilizando "(*host)[len(*host)-1]" o go primeiro desreferencia o ponteiro e depois acessa o valor do indice, por causa dos parenteses o go entende que deve priorizar o que está dentro deles, destá forma o go entende que deve desreferenciar o ponteiro e depois acessar o valor do indice do slice
}