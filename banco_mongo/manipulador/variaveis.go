package manipulador

import "html/template"

//ModeloOla armazena os modelos de página que serão executados pelo manipulador de ola
var ModeloOla = template.Must(template.ParseFiles("html/ola.html"))

//ModeloLocal armazena os modelos de página que serão executados pelo manipulador de local
var ModeloLocal = template.Must(template.ParseFiles("html/local.html"))
