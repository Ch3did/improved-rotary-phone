package main

import (
	"fmt"

	"github.com/gorilla/websocket"
)

type Usuario struct {
	nome string
	id   *websocket.Conn
}

type Mensagem struct {
	Usuario *Usuario
	text    string
}

func main() {
	n := PrimeiraAcao()
	fmt.Println(n)
}

//PrimeiraAcao é o método de abertura da API
func PrimeiraAcao() string {
	var nome string
	fmt.Println("----------Chat da Uolminal----------")
	fmt.Println("Digite seu apelido: ")
	fmt.Scanln(&nome)
	return nome
}

func WSocket() {

}
