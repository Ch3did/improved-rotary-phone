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

func mantendoConexao(conn *websocket.Conn) {
	for {
		tipoMensagem, p, err := conn.ReadMessage()
		if err != nil {
			log.Println("mantendo a conexão")
			return
		}
		fmt.Println(string(p))

		if err := conn.WriteMessage(tipoMensagem, p); err != nil {
			log.Println("a msg")
			return
		}
	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("endpoint")
	}

	log.Println("Conexão bem sucedida...")

	mantendoConexao(ws)
}

//Routes implementa as rotas da aplicação WEB
func Routes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", wsEndpoint)
}

func main() {
	fmt.Println("Hello World")
	Routes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}

//Ajustar as msg e recebimentos
