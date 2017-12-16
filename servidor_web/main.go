package main

import (
	"fmt"
	"net/http"
	"servidor_web/manipulador"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Olá mundo")
	})

	http.HandleFunc("/funcao", manipulador.Funcao)
	http.HandleFunc("/ola", manipulador.Ola)

	fmt.Println("Hey...listen...")
	http.ListenAndServe(":9000", nil)
}
