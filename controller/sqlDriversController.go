package controller

import (
	"github.com/ylanzinhoy/sollievo/internal/model"
)

const (
	// names

	postgres = "postgres"
	mySQL    = "mysql"
	mongoDB  = "mongodb"

	// packages
	postgresPackage = "go get github.com/lib/pq"
	mysqlPackage    = "go get github.com/go-sql-driver/mysql"
	mongoPackage    = "go get go.mongodb.org/mongo-driver/mongo"
)

func SqlDriversController() {

	model.NewControllerModel("sqlDrivers",
		map[string]string{
			mongoDB:  mongoPackage,
			postgres: postgresPackage,
			mySQL:    mysqlPackage,
		},
		&commandsStruct).ProcessCommand()
}
