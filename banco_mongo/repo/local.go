package repo

import (
	"banco_mongo/model"

	"gopkg.in/mgo.v2/bson"
)

//ObtemLocal obtem o local no MongoDB
func ObtemLocal(codigoTelefone int) (local model.Local, err error) {
	sessao := SessaoMongo.Copy()
	defer sessao.Close()
	colecao := sessao.DB("cursodego").C("local")
	err = colecao.Find(bson.M{"telcode": codigoTelefone}).One(&local)
	return
}

//SalvaLog registra a consulta ao local
func SalvaLog(reg model.RegistroLog) (err error) {
	sessao := SessaoMongo.Copy()
	defer sessao.Close()
	colecao := sessao.DB("cursodego").C("local")
	colecao.Insert(reg)
	return
}
