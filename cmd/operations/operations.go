package checks

import (
	checkDB "Synapse/cmd/operations/check/db"
	createDB "Synapse/cmd/operations/create/db"
	dropDB "Synapse/cmd/operations/drop/db"
	HelperCmd "Synapse/cmd/operations/help"
)

type Check interface {
	Run()
}

var checkMap = map[string]Check{
	"--help":      HelperCmd.HelperCmd{},
	"--check-db":  checkDB.CheckDatabase{},
	"--create-db": createDB.DatabaseCreator{},
	"--drop-db":   dropDB.DatabaseDropper{},
}

// Get retorna o check se ele existir no mapa
func Get(flag string) (Check, bool) {
	val, ok := checkMap[flag]
	return val, ok
}
