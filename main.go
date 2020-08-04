package main

import (
	"fmt"
	"log"

	"github.com/Ch3did/improved-rotary-phone/client"
	"github.com/Ch3did/improved-rotary-phone/server"
	"github.com/Ch3did/improved-rotary-phone/utils"
)

func main() {
	fmt.Println("CHAT UOLDACHINA")
	fmt.Println("Deseja iniciar o server (1) ou startar o client (2)?...")
	escolha := utils.EscolhaDeUso()
	if escolha == 1 {
		server.Inicio()
	}
	if escolha == 2 {
		client.Inicio()
	} else {
		log.Fatal("Opção inválida")
	}

}
