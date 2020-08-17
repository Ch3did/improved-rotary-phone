package utils

import (
	"bufio"
	"os"
)

//Scan recebe o input do usuário para escoler o tipo da aplicação
func Scan() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	return text
}
