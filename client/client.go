package client

import (
	"bufio"
	"fmt"
	"os"

	"github.com/gorilla/websocket"
	"main.go/utils"
)

// Inicio inicia
func Inicio(nick string) {
	conn, _, err := websocket.DefaultDialer.Dial("localhost:3000/", nil)
	if err != nil {

	}
	defer conn.Close()

	go func() {
		for {
			// Lê a msg
			_, msg, err := conn.ReadMessage()
			if err != nil {
				break
			}
			s := string(msg)
			fmt.Println(s)
		}
	}()

	// Envia a msg
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		msg := scanner.Text()
		msgF := utils.FormatMSG(nick, msg)
		msgFB := []byte(msgF)
		err = conn.WriteMessage(websocket.TextMessage, msgFB)
	}
}
