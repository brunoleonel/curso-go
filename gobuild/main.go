package main

import (
	"encoding/json"
	"fmt"

	"github.com/gobuild/model"
)

/*
GOOS=windows GOARCH=386 go build -o meuappwindows.exe
GOOS=linux GOARCH=arm go build -o meuappraspberry -v github.com/gobuild/
*/
func main() {

	casa := model.Imovel{}
	casa.Nome = "Casa amarela"
	casa.X = 18
	casa.Y = 25
	casa.SetValor(60000)

	fmt.Printf("Casa é %+v\r\n", casa)
	fmt.Printf("O valor da casa é %d\r\n", casa.GetValor())
	objJSON, _ := json.Marshal(casa)
	fmt.Println("A casa em JSON: ", string(objJSON))
}
