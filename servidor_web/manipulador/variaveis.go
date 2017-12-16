package manipulador

import "html/template"

//Modelos armazena os modelos de página que serão executados pelos manipuladores
var Modelos = template.Must(template.ParseFiles("html/ola.html"))
