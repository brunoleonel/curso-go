package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"json_unmarshall/model"
	"net/http"
	"time"
)

func main() {

	cliente := &http.Client{
		Timeout: time.Second * 30,
	}

	request, err := http.NewRequest("GET", "https://jsonplaceholder.typicode.com/posts/1", nil)
	request.SetBasicAuth("teste", "teste")
	if err != nil {
		fmt.Println("[main] Erro ao criar o request", err.Error())
		return
	}
	resposta, err := cliente.Do(request)
	if err != nil {
		fmt.Println("[main] Erro ao ler o json da resposta", err.Error())
		return
	}
	defer resposta.Body.Close()

	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[main] Erro ao ler o json da resposta", err.Error())
			return
		}
		fmt.Println(" ")
		post := model.BlogPost{}
		err = json.Unmarshal(corpo, &post)
		if err != nil {
			fmt.Println("[main] Erro ao converter o json da resposta", err.Error())
			return
		}
		fmt.Printf("Post: %+v\r\n", post)
	}

}
