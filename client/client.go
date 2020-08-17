package client

import (
	"fmt"

	"github.com/Ch3did/improved-rotary-phone/utils"
)

var (
	msg  string
	nick string
)

//Inicio inicializa o pacote
func Inicio() {
	fmt.Println("Me dê o seu nick: ")
	nick = utils.Scan()
	mandarMsg()
}

func mandarMsg() {
	for {
		msg = utils.Scan()
		fmt.Printf("[%s] %s", nick, msg)
	}
}
