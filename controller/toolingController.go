package controller

import (
	"fmt"
	"github.ylanzinhoy.tooling_golang/enums"
	"github.ylanzinhoy.tooling_golang/model"
	"github.ylanzinhoy.tooling_golang/service"
)

func ToolingController() {
	toolsChoices := make([]string, 0, 2)
	commandsStruct := service.CommandsStruct{}
	modelTools := model.Tools{}

	tools := commandsStruct.Choices(modelTools.ToolsChoice(append(toolsChoices,
		enums.Gorm, enums.Viper, enums.Wire, enums.Default)), "tools")

	for _, choice := range tools {
		switch choice {
		case enums.Gorm:
			fmt.Println("Testando gorm")
		case enums.Viper:
			fmt.Println("Testando Viper")

		case enums.Default:
			fmt.Println("Saindo...")
			break
		}

	}
}
