//;;package main

import (
	"fmt"

	"github.com/gorilla/websocket"
)

//Usuario ...
type Usuario struct {
	nome string
	id   *websocket.Conn
}

//Mensagem ...
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


.//HTTP server with WebSocket endpoint
func Server() {
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    ws, err := NewHandler(w, r)
    if err != nil {
         // handle error
    }
    if err = ws.Handshake(); err != nil {
        // handle error
    }
}