package main

import "fmt"

//Variável publica
//Endereco - Exportar com a variável documentada. Caso contrário causa warning
var Endereco string

//Variável privada
var telefone = "9999-9999"

//Variáveis globais
var quantidade int
var comprou bool

//Variáveis globais
var (
	valor float64
	palavras rune
)

func main() {
	//variavel de escopo de bloco
	comando := "Hadougui"
	fmt.Printf("Endereço: %s \r\n", Endereco)
	fmt.Printf("Quantidade: %d \r\n", quantidade)
	fmt.Printf("Valor: %f \r\n", valor)

	fmt.Printf("Solta %s Ryu", comando)
}