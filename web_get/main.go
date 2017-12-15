package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {

	cliente := &http.Client{
		Timeout: time.Second * 30,
	}
	resposta, err := cliente.Get("https://www.google.com.br")
	if err != nil {
		fmt.Println("[main] Erro ao abrir a página do google no Brasil: ", err.Error())
		return
	}
	defer resposta.Body.Close()

	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[main] Erro ao ler o conteudo da página do Google Brasil", err.Error())
			return
		}
		fmt.Println(string(corpo))
	}

	request, err := http.NewRequest("GET", "https://www.google.com.br", nil)
	request.SetBasicAuth("teste", "teste")
	if err != nil {
		fmt.Println("[main] Erro ao criar o request", err.Error())
		return
	}
	resposta, err = cliente.Do(request)
	if err != nil {
		fmt.Println("[main] Erro ao ler o conteudo da página do Google Brasil", err.Error())
		return
	}
	fmt.Println(resposta)
}
