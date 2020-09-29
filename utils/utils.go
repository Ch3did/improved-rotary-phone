package utils

import (
	"bufio"
	"os"
)

var msgFormatado string

//Scan recebe o input do usuário
func Scan() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	return text
}

func FormatMSG(nick, msg string) (msgFormatado string) {
	msgFormatado = ("[" + nick + "]: " + msg + "\n")
	return

}
