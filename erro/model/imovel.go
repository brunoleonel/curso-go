package model

import "errors"

//Imovel representa informações de um imóvel
type Imovel struct {
	X     int    `json:"coordenada_x"`
	Y     int    `json:"coordenada_y"`
	Nome  string `json:"nome"`
	valor int
}

//ErrValorDeImovelAbaixo é um erro lançado caso o valor atribuído ao imóvel seja muito baixo
var ErrValorDeImovelAbaixo = errors.New("O valor informado é inválido por ser muito baixo")

//ErrValorDeImovelAcima é um erro lançado caso o valor atribuído ao imóvel seja muito alto
var ErrValorDeImovelAcima = errors.New("O valor informado é inválido por ser muito alto")

//SetValor atribui o valor ao imóvel
func (i *Imovel) SetValor(valor int) (err error) {
	err = nil
	if valor < 50000 {
		err = ErrValorDeImovelAbaixo
		return
	} else if valor > 100000 {
		err = ErrValorDeImovelAcima
	}
	i.valor = valor
	return
}

//GetValor retorna o valor do imóvel
func (i *Imovel) GetValor() int {

	return i.valor
}
