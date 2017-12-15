package main

import (
	"fmt"
)

func main() {
	capitais := []string{"Lisboa"}
	//fmt.Println(capitais, len(capitais), cap(capitais))
	capitais = append(capitais, "Brasília")
	//fmt.Println(capitais, len(capitais), cap(capitais))
	capitais[1] = "Brasilia"
	//fmt.Println(capitais, len(capitais), cap(capitais))
	cidades := make([]string, 5)
	//fmt.Println(cidades, len(cidades), cap(cidades))
	cidades[0] = "Nova York"
	cidades[1] = "Londres"
	cidades[2] = "Madeira"
	cidades[3] = "Tokio"
	cidades[4] = "Singapura"
	fmt.Println(cidades, len(cidades), cap(cidades))
	for indice, cidade := range cidades {
		fmt.Printf("Cidade[%d] = %s\r\n", indice, cidade)
	}
	capitaisAsia := cidades[3:5]
	fmt.Println(capitaisAsia)
	temp1 := cidades[:2]
	fmt.Println(temp1)
	temp2 := cidades[len(cidades)-2:]
	fmt.Println(temp2)
	indiceDoItemARetirar := 2
	temp := cidades[:indiceDoItemARetirar]
	temp = append(temp, cidades[indiceDoItemARetirar:]...)
	copy(cidades, temp)
	fmt.Println(cidades)
}
