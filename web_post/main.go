package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"web_post/model"
)

func main() {

	cliente := &http.Client{
		Timeout: time.Second * 30,
	}

	usuario := model.Usuario{}
	usuario.ID = 1
	usuario.Nome = "Jeff Prestes"

	conteudoEnviar, err := json.Marshal(usuario)
	if err != nil {
		fmt.Println("[main] Erro ao criar o json", err.Error())
		return
	}

	request, err := http.NewRequest("POST", "https://requestb.in/pgwbsdpg", bytes.NewBuffer(conteudoEnviar))
	request.SetBasicAuth("fizz", "buzz")
	request.Header.Set("content-type", "application/json; charset=utf-8")
	if err != nil {
		fmt.Println("[main] Erro ao enviar o request", err.Error())
		return
	}
	resposta, err := cliente.Do(request)
	if err != nil {
		fmt.Println("[main] Erro ao ler o conteudo da resposta", err.Error())
		return
	}
	fmt.Println(resposta.Status)
}
