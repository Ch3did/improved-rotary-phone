package main

import (
	"fmt"
	"log"

	"main.go/client"
	"main.go/server"
	"main.go/utils"
)

func main() {
	fmt.Println("CHAT UOLDACHINA")
	fmt.Println("Deseja iniciar o server (1) ou startar o client (2)?...")
	escolha := utils.Scan()
	if escolha == "1" {
		server.Inicio()
	}
	if escolha == "2" {
		fmt.Print("Me dê o seu nick: ")
		nick := utils.Scan()
		client.Inicio(nick)
	} else {
		log.Fatal("Opção Inválida ")
	}
}
