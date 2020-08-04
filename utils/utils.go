package utils

import (
	"bufio"
	"fmt"
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
	} else {
		fmt.Println("Erro...")
		return 0
	}

}
