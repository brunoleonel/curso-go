package main

import (
	"encoding/json"
	"erro/model"
	"fmt"
)

func main() {

	casa := model.Imovel{}
	casa.Nome = "Casa amarela"
	casa.X = 18
	casa.Y = 25
	if err := casa.SetValor(120000); err != nil {
		fmt.Println("[main] Houve um erro na atribuição do valor da casa: ", err.Error())
		if err == model.ErrValorDeImovelAcima {
			fmt.Println("Corretor, por favor verifique sua avaliação")
		}
		return
	}

	fmt.Printf("Casa é %+v\r\n", casa)
	fmt.Printf("O valor da casa é %d\r\n", casa.GetValor())
	objJSON, err := json.Marshal(casa)

	if err != nil {
		fmt.Println("[main] Houve um erro na geração do JSON: ", err.Error())
		return
	}

	fmt.Println("A casa em JSON: ", string(objJSON))
}
