package app

import (
	"fmt"
	"linha-de-comando/utils"
	"net"
	"github.com/urfave/cli"
)

// Gerar retorna a aplicação de linha de comando pronta para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicativo de Linha de Comando"
	app.Usage = "Busca IPs e Nomes de Servidor na Internet"
	app.Version = "1.0.0"

	flags := []cli.Flag{
		cli.StringFlag{
			Name: "host",
			Usage: "O host que será verificado",
			Required: true,
			Value: "google.com", // Valor padrão
		},
	}

	app.Commands = []cli.Command{
		// Comando de IP
		{
			Name:  "ip",
			Usage: "Busca IPs de endereços na internet",
			Flags: []cli.Flag{ // Está aqui só para mostrar como ter flags diferentes para cada comando, mas poderia ser a mesma coisa que a variável flags
				cli.StringFlag{
					Name: "host",
					Usage: "O host que será verificado",
					Required: true,
					Value: "google.com", // Valor padrão
				},
			},
			Action: buscarIps,
		},

		// Comando de Nome de Servidor
		{
			Name: "server-names",
			Usage: "Busca o nome dos servidores na internet",
			Flags: flags, // Utiliza a variável flags que foi criada anteriormente
			Action: buscarNomesDosServidores,
		},
	}

	return app
}


func buscarIps(c *cli.Context) error {
	host := c.String("host")
	
	utils.FormatarEnderecoHost(&host) // altera diretamente o valor da variável host no endereço de memória

	fmt.Printf("Buscando IPs do host: %s\n", host)

	ips, error := net.LookupIP(host)
	if error != nil {
		return error
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}

	return nil
}

func buscarNomesDosServidores(c *cli.Context) error {
	host := c.String("host")

	utils.FormatarEnderecoHost(&host)

	nomesServidores, erro := net.LookupNS(host) // NS = Name Server
	if erro != nil {
		return erro
	}

	for _, nomeSevidor := range nomesServidores {
		fmt.Println(nomeSevidor)
	}

	return nil
}