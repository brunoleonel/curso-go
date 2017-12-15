package main

import (
	"fmt"
	"interfaces/model"
)

func main() {

	jojo := model.Ave{}
	jojo.Nome = "Jojo da Silva"

	queroAcordarComUmCacarejo(jojo)
	queroOuvirUmPatoNoLago(jojo)
}

func queroAcordarComUmCacarejo(g model.Galinha) {
	fmt.Println(g.Cacareja())
}

func queroOuvirUmPatoNoLago(p model.Pata) {
	fmt.Println(p.Quack())
}
