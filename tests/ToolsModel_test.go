package tests

import (
	"testing"

	"github.ylanzinhoy.tooling_golang/model"
)

func TestToolsModel(t *testing.T) {
	toolsModel := &model.Tools{
		Tools: map[string]string{
			"mongoDB":   "mongoPackage",
			"cassandra": "cassandraPackage",
			"postgres":  "postgresPackage",
			"mySQL":     "mysqlPackage",
		},
	}

	res := toolsModel.ToolsChoice()
	if res == nil {
		t.Errorf("Array empty %v", res)
	}

	


}
