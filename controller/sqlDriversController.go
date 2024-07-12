package controller

import "github.ylanzinhoy.tooling_golang/enums"

func (t *ToolingControllerUpper) SqlDriversController() {

	driversChoice := make([]string, 0)

	tools := commandsStruct.Choices(modelTools.ToolsChoice(append(driversChoice, enums.Postgres, enums.MySQL)),
		"sqlDrivers")

	for _, choice := range tools {
		switch choice {
		case enums.Postgres:
			t.Exec(&commandsStruct, enums.Postgres, enums.PostgresPackage)
		case enums.MySQL:
			t.Exec(&commandsStruct, enums.MySQL, enums.MysqlPackage)
		}
	}
}
