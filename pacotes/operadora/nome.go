package operadora

import (
	"pacotes/prefixo"
	"strconv"
)

//NomeOperadora representa o nome da operadora
var NomeOperadora = "XPTO Telecom"

//PrefixoDaCapitalOperadora é o prefixo + nome da operadora
var PrefixoDaCapitalOperadora = strconv.Itoa(prefixo.Capital) + " " + NomeOperadora
