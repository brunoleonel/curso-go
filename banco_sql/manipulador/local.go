package manipulador

import (
	"banco_sql/model"
	"banco_sql/repo"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

//Local é o manipulador da rota /local
func Local(w http.ResponseWriter, r *http.Request) {
	local := model.Local{}

	codigoTelefone, err := strconv.Atoi(r.URL.Path[7:])
	if err != nil {
		http.Error(w, "Não foi enviado um número válido", http.StatusBadRequest)
		fmt.Println("[Local] Erro: ", err.Error())
		return
	}

	sql := "select country, city, telcode from place where telcode = ?"
	linha, err := repo.Db.Queryx(sql, codigoTelefone)
	if err != nil {
		http.Error(w, "Não foi possível executar a query", http.StatusBadRequest)
		fmt.Println("[Local] Erro: ", err.Error())
		return
	}

	for linha.Next() {
		err = linha.StructScan(&local)
		if err != nil {
			http.Error(w, "Não foi possível fazer o binding do resultado", http.StatusBadRequest)
			fmt.Println("[Local] Erro: ", err.Error())
			return
		}
	}

	if err := ModeloLocal.ExecuteTemplate(w, "local.html", local); err != nil {
		http.Error(w, "Houve um erro na renderização da página", http.StatusInternalServerError)
		fmt.Println("[Local] Erro na execução do modelo: ", err.Error())
		return
	}

	sql = "insert into logquery(daterequest) values(?)"
	resultado, err := repo.Db.Exec(sql, time.Now().Format("02/01/2006 15:04:05"))
	if err != nil {
		fmt.Println("[Local] Erro ao inserir o log: ", sql, " - ", err.Error())
		return
	}
	linhasAfetadas, err := resultado.RowsAffected()
	if err != nil {
		fmt.Println("[Local] Erro ao coletar número de linhas afetadas: ", sql, " - ", err.Error())
		return
	}
	fmt.Println("Sucesso! ", linhasAfetadas, " linhas afetadas.")
}
