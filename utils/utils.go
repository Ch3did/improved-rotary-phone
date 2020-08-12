package utils

import (
	"bufio"
	"net/http"
	"os"
)

//EscolhaDeUso recebe o input do usuário para escoler o tipo da aplicação
func EscolhaDeUso() int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	// implementar tratamento da string recebida
	if text == "1" || text == "server" {
		return 1
	}
	// implementar tratamento da string recebida
	if text == "2" || text == "client" {
		return 1
	}
	return 0
}

//Routes implementa as rotas da aplicação WEB
func Routes() {
	http.HandleFunc("/ws", wsEndpoint)
}
