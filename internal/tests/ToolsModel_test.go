package tests

import (
	"testing"

	"github.com/ylanzinhoy/sollievo/internal/model"
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

	_, err := toolsModel.ToolsChoice()
	if err != nil {
		t.Errorf("Array error %v", err)
	}

}
