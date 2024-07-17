package controller

import (
	"github.ylanzinhoy.tooling_golang/enums"
)

const (
	// names

	postgres  = "postgres"
	mySQL     = "mysql"
	cassandra = "cassandra"
	mongoDB   = "mongodb"

	// packages
	postgresPackage  = "go get github.com/lib/pq"
	mysqlPackage     = "go get -u github.com/go-sql-driver/mysql"
	cassandraPackage = "go get github.com/gocql/gocql"
	mongoPackage     = "go get go.mongodb.org/mongo-driver/mongo"
)

func (t *ToolingControllerUpper) SqlDriversController() {

	driversChoice := make([]string, 0)

	tools := commandsStruct.Choices(modelTools.ToolsChoice(append(driversChoice, postgres,
		mySQL, cassandra, mongoDB, enums.Default, enums.Back)),
		"sqlDrivers")

	maps := map[string]string{
		mongoDB:   mongoPackage,
		cassandra: cassandraPackage,
		postgres:  postgresPackage,
		mySQL:     mysqlPackage,
	}

	t.executeChoices(maps, tools)
}
