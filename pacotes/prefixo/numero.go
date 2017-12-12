package prefixo

import "strconv"

//Capital representa o número do prefixo de telefone para a capital
var Capital = 11

//Essa é privada, se for tentar usar usar direto causa erro. Se for o caso precisa ser exportada -> Teste
var teste = "teste"

//TesteComPrefixo é um teste
var TesteComPrefixo = teste + " " + strconv.Itoa(Capital)
