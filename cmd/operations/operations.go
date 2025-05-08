package checks

import (
	checkDB "Synapse/cmd/operations/check/db"
	createDB "Synapse/cmd/operations/create/db"
)

type Check interface {
	Run()
}

var checkMap = map[string]Check{
	"--check-db":  checkDB.CheckDatabase{},
	"--create-db": createDB.DatabaseCreator{},
}

// Get retorna o check se ele existir no mapa
func Get(flag string) (Check, bool) {
	val, ok := checkMap[flag]
	return val, ok
}
