package main

import (
	"fmt"
)

//Imovel é um struct que armazena dados do imóvel
type Imovel struct {
	X     int
	Y     int
	Nome  string
	Valor int
}

func main() {

	casa := Imovel{}
	fmt.Printf("A casa é %+v\r\n", casa)

	apartamento := Imovel{17, 45, "Meu AP", 760000}
	fmt.Printf("O apartamento é %+v\r\n", apartamento)

	chacara := Imovel{
		Y:     85,
		Nome:  "Chacara",
		Valor: 55,
		X:     22,
	}
	fmt.Printf("A chacara é %+v\r\n", chacara)

	casa.Nome = "Lar doce lar"
	casa.Valor = 450000
	casa.X = 18
	casa.Y = 25
	fmt.Printf("A casa é %+v\r\n", casa)
}
