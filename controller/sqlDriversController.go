package controller

import (
	"github.ylanzinhoy.tooling_golang/model"
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
	maps := map[string]string{
		mongoDB:   mongoPackage,
		cassandra: cassandraPackage,
		postgres:  postgresPackage,
		mySQL:     mysqlPackage,
	}

	toolsStruct := &model.Tools{
		Tools: maps,
	}

	tools := commandsStruct.Choices(toolsStruct.ToolsChoice(), "sqlDrivers")

	t.executeChoices(maps, tools)
}
