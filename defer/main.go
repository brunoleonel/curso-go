package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"escreverarquivos/model"
	"fmt"
	"os"
	"strings"
)

func main() {
	arquivo, err := os.Open("cidades.csv")
	if err != nil {
		fmt.Println("[main] Houve um erro ao abrir o arquivo: ", err.Error())
		return
	}
	defer arquivo.Close()

	// scanner := bufio.NewScanner(arquivo)
	// for scanner.Scan() {
	// 	linha := scanner.Text()
	// 	fmt.Println("O conteúdo da linha é: ", linha)
	// }

	leitorCsv := csv.NewReader(arquivo)
	conteudo, err := leitorCsv.ReadAll()
	if err != nil {
		fmt.Println("[main] Houve um erro ao ler o arquivo com leitor CSV", err.Error())
		return
	}

	arquivoJSON, err := os.Create("cidades.json")
	if err != nil {
		fmt.Println("[main] Houve um erro ao criar o arquivo JSON", err.Error())
		return
	}
	defer arquivoJSON.Close()

	escritor := bufio.NewWriter(arquivoJSON)
	escritor.WriteString("[\r\n")

	for _, linha := range conteudo {
		for indiceItem, item := range linha {
			dados := strings.Split(item, "/")
			cidade := model.Cidade{}
			cidade.Nome = dados[0]
			cidade.Estado = dados[1]
			fmt.Printf("Cidade: %+v\r\n", cidade)
			cidadeJSON, err := json.Marshal(cidade)
			if err != nil {
				fmt.Println("[main] Houve um erro ao criar o arquivo JSON: ", item, " - ", err.Error())
			}
			escritor.WriteString(" " + string(cidadeJSON))
			if (indiceItem + 1) < len(linha) {
				escritor.WriteString(",\r\n")
			}
		}
	}
	escritor.WriteString("\r\n]")
	escritor.Flush()
}
