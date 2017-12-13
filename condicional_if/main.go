package main

import (
	"fmt"
)

func main() {

	meses := 0
	situacao := true
	cidade := "Teste"

	//< > != == <= >= && ||
	if meses <= 6 {
		fmt.Println("Esse credor deve há pouco tempo.")
	}

	if situacao {
		fmt.Println("Ele está devendo")
	}

	if !situacao {
		fmt.Println("Ele não está devendo")
	}

	if cidade == "Teste" {
		fmt.Println("O cliente vive na cidade Teste")
	}

	if descricao, status := haQuantoTempoEstaDevendo(meses); status {
		fmt.Println("Qual a situação do cliente? ", descricao)
		return
	}

	//A instrução abaixo não compila - a variavel descricao não funciona fora do escopo do if
	//fmt.Println("Me passa o status?", descricao)

	fmt.Println("Obrigado por nos consultar.")
}

func haQuantoTempoEstaDevendo(meses int) (descricao string, status bool) {

	if meses > 0 {
		status = true
		descricao = "O cliente está devendo."
		return
	}

	descricao = "O cliente está com o carnê em dia"
	return
}
