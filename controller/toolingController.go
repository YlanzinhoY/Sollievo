package controller

import (
	"fmt"
	"github.ylanzinhoy.tooling_golang/enums"
	"github.ylanzinhoy.tooling_golang/model"
	"github.ylanzinhoy.tooling_golang/service"
	"net/http"
)

func ToolingController() {
	toolsChoices := make([]string, 0, 3)
	commandsStruct := service.CommandsStruct{}
	modelTools := model.Tools{}

	tools := commandsStruct.Choices(modelTools.ToolsChoice(append(toolsChoices,
		enums.Gorm, enums.Viper, enums.Wire, enums.Default)), "tools")

	//region switch case
	for _, choice := range tools {
		switch choice {
		case enums.Gorm:
			err := commandsStruct.CommandRunner(enums.Gorm, enums.GormPackage)
			if err != nil {
				fmt.Println(http.StatusBadRequest)
				break
			}
			fmt.Printf("%d\n", http.StatusOK)
		case enums.Viper:
			err := commandsStruct.CommandRunner(enums.Viper, enums.ViperPackage)
			if err != nil {
				fmt.Println(http.StatusBadRequest)
				break
			}
			fmt.Printf("%d\n", http.StatusOK)
		case enums.Wire:
			err := commandsStruct.CommandRunner(enums.Wire, enums.WirePackage)
			if err != nil {
				fmt.Println(http.StatusBadRequest)
				break
			}
			fmt.Printf("%d\n", http.StatusOK)
		case enums.Default:
			fmt.Println("Saindo...")
			break
		}
	}

	//endregion
}
