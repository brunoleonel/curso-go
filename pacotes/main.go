package main

import (
	"fmt"
	"pacotes/operadora"
	"pacotes/prefixo"
)

//NomeDoUsuario é o nome do usuário do sistema
var NomeDoUsuario = "Jeff"

func main() {
	fmt.Printf("Nome do usuário: %s\r\n", NomeDoUsuario)
	fmt.Printf("Prefixo da capital: %d\r\n", prefixo.Capital)
	fmt.Printf("Nome da operadora: %s\r\n", operadora.NomeOperadora)
	fmt.Printf("Teste com prefixo: %s\r\n", prefixo.TesteComPrefixo)
}
