package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

//Porta ...
var Porta = ":8080"

func mantendoConexao(conn *websocket.Conn) {
	for {
		tipoMensagem, msg, _ := conn.ReadMessage()
		conn.WriteMessage(tipoMensagem, msg)
	}
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, _ := upgrader.Upgrade(w, r, nil)

	log.Println("Conexão bem sucedida...")

	mantendoConexao(ws)
}

//Routes implementa as rotas da aplicação WEB
func Routes() {
	http.HandleFunc("/ws", wsEndpoint)
}

//Inicio inicializa o pacote
func Inicio() {
	fmt.Println("Hello World")
	Routes()
	log.Fatal(http.ListenAndServe(Porta, nil))
}

//Ajustar as msg e recebimentos
