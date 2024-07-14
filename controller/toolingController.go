package controller

import (
	"github.ylanzinhoy.tooling_golang/enums"
)

func (t *ToolingControllerUpper) ToolingController() {
	toolsChoices := make([]string, 0)

	tools := commandsStruct.Choices(modelTools.ToolsChoice(append(toolsChoices,
		enums.Gorm, enums.Viper, enums.Wire, enums.Default, enums.Back)), "tools")

	for _, choice := range tools {
		switch choice {
		case enums.Gorm:
			t.Exec(&commandsStruct, enums.Gorm, enums.GormPackage)
		case enums.Viper:
			t.Exec(&commandsStruct, enums.Viper, enums.ViperPackage)
		case enums.Wire:
			t.Exec(&commandsStruct, enums.Wire, enums.WirePackage)
		case enums.Prometheus:
			t.Exec(&commandsStruct, enums.Prometheus, enums.PrometheusPackage)
		case enums.Default:
			t.Exit(&commandsStruct)
		case enums.Back:
			t.Back(&commandsStruct, "tooling_golang")
		default:
			return
		}
	}
}
