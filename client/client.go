package client

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/Ch3did/improved-rotary-phone/utils"
	"github.com/sacOO7/gowebsocket"
)

var (
	msg          string
	nick         string
	msgFormatado string
)

//FormatOfMsg formata o jeito de como a msg é usada no terminal
func FormatOfMsg() (msgFormatado string) {
	fmt.Println("Me dê o seu nick: ")
	nick = utils.Scan()
	msg = utils.Scan()
	msgFormatado = "[" + nick + "]" + msg + "\n"
	return
}

//Inicio inicia o clinte
func Inicio() {
	for {
		interrupt := make(chan os.Signal, 1)
		signal.Notify(interrupt, os.Interrupt)
		socket := gowebsocket.New("ws://localhost:8080/ws")

		socket.OnConnectError = func(err error, socket gowebsocket.Socket) {
			log.Fatal("Received connect error - ", err)
		}

		socket.OnConnected = func(socket gowebsocket.Socket) {
			log.Println("Connected to server")
		}

		socket.OnTextMessage = func(msg string, socket gowebsocket.Socket) {
			log.Println("Received message - " + msg)
		}

		socket.OnPingReceived = func(data string, socket gowebsocket.Socket) {
			log.Println("Received ping - " + data)
		}

		socket.OnPongReceived = func(data string, socket gowebsocket.Socket) {
			log.Println("Received pong - " + data)
		}
		socket.OnDisconnected = func(err error, socket gowebsocket.Socket) {
			log.Println("Disconnected from server ")
			return
		}

		socket.Connect()
		send := FormatOfMsg()
		socket.SendText(send)
		for {
			select {
			case <-interrupt:
				log.Println("interrupt")
				socket.Close()
				return
			}
		}

	}

}
