package manipulador

import "net/http"
import "fmt"

//Funcao é uma função que atende um manipulador
func Funcao(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Essa é uma função num pacote que atende um manipulador")
}
