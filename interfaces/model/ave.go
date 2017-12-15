package model

//Galinha é a interface que define as funções de uma galinha
type Galinha interface {
	Cacareja() string
}

//Pata é a interface que define as funções de uma pata
type Pata interface {
	Quack() string
}

//Ave é um tipo que representa aves
type Ave struct {
	Nome string
}

//Cacareja é a função de cacarejo
func (a Ave) Cacareja() string {
	return "Cocóricó"
}

//Quack é a função de quaquejar
func (a Ave) Quack() string {
	return "Quack"
}
