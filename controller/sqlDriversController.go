package controller

import (
	"github.com/ylanzinhoy/sollievo/internal/model"
)

const (
	// names

	postgres  = "postgres"
	mySQL     = "mysql"
	cassandra = "cassandra"
	mongoDB   = "mongodb"

	// packages
	postgresPackage  = "go get github.com/lib/pq"
	mysqlPackage     = "go get github.com/go-sql-driver/mysql"
	cassandraPackage = "go get github.com/gocql/gocql"
	mongoPackage     = "go get go.mongodb.org/mongo-driver/mongo"
)

func SqlDriversController() {

	model.NewControllerModel("sqlDrivers",
		map[string]string{
			mongoDB:   mongoPackage,
			cassandra: cassandraPackage,
			postgres:  postgresPackage,
			mySQL:     mysqlPackage,
		},
		&commandsStruct).ProcessCommand()
}
