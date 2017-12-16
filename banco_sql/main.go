package main

import (
	"banco_sql/manipulador"
	"banco_sql/repo"
	"fmt"
	"net/http"
)

func init() {
	fmt.Println("Subindo o servidor...")
}

func main() {

	err := repo.AbreConexaoComBancoDeDadosSQL()
	if err != nil {
		fmt.Println("Houve um erro ao conectar com o banco de dados: ", err.Error())
		return
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Olá mundo")
	})

	http.HandleFunc("/funcao", manipulador.Funcao)
	http.HandleFunc("/ola", manipulador.Ola)
	http.HandleFunc("/local/", manipulador.Local)

	fmt.Println("Hey...listen...")
	http.ListenAndServe(":9000", nil)
}
