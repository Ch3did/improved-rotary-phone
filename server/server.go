package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

// Inicio inicia
func Inicio() {

	http.HandleFunc("/", handler)

	http.ListenAndServe(":3000", nil)
}

func handler(writer http.ResponseWriter, request *http.Request) {
	conn, err := upgrader.Upgrade(writer, request, nil)
	if err != nil {
		fmt.Println(err)
	}

	for {
		// Lê a msg
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		// joga no console

		fmt.Println(string(msg))

		// Devolvendo a mensagem recebida
		conn.WriteMessage(msgType, msg)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
