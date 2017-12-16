package repo

import (
	"gopkg.in/mgo.v2"
)

//SessaoMongo armazena os dados da sessão do mongodb
var SessaoMongo *mgo.Session

//AbreSessaoComMongo abre a sessão com o mongodb
func AbreSessaoComMongo() (err error) {
	err = nil
	SessaoMongo, err = mgo.Dial("mongodb://localhost:27017/cursodego")
	return
}
