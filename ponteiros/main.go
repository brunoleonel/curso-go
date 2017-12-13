package main

import "fmt"

//Imovel é uma struct que armazena dados de um imóvel
type Imovel struct {
	X     int
	Y     int
	Nome  string
	Valor int
}

func main() {

	casa := new(Imovel)
	fmt.Printf("Casa é: %p - %+v\r\n", &casa, casa)

	chacara := Imovel{17, 28, "Chacara linda", 280000}
	fmt.Printf("Chacara é: %p - %+v\r\n", &chacara, chacara)

	mudaImovel(&chacara)
	fmt.Printf("Chacara é: %p - %+v\r\n", &chacara, chacara)
}

func mudaImovel(imovel *Imovel) {
	imovel.X = imovel.X + 10
	imovel.Y = imovel.Y - 5
}
