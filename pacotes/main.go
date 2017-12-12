package main

import "fmt"
import "github.com/pacotes/prefixo"

//Nome do usuário é o nome do usuário do sistema
var NomeDoUsuario = "Jeff"

func main() {
	fmt.Printf("Nome do usuário: %s\r\n", NomeDoUsuario)
	fmt.Printf("Prefixo da capital: %d\r\n", prefixo.Capital)
}