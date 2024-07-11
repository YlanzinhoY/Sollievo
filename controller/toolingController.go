package controller

import (
	"fmt"
	"github.ylanzinhoy.tooling_golang/enums"
	"github.ylanzinhoy.tooling_golang/model"
	"github.ylanzinhoy.tooling_golang/service"
	"net/http"
)

type ToolsServiceInterface interface {
	Exec(commandsStruct *service.CommandsStruct, value, goPackage string)
}

type ToolingControllerUpper struct{}

func (t *ToolingControllerUpper) Exec(commandsStruct *service.CommandsStruct, value, goPackage string) {
	err := commandsStruct.CommandRunner(value, goPackage)
	if err != nil {
		fmt.Println(http.StatusBadRequest)
	}
	fmt.Printf("%d\n", http.StatusOK)
}

func (t *ToolingControllerUpper) ToolingController() {
	toolsChoices := make([]string, 0, 3)
	commandsStruct := service.CommandsStruct{}
	modelTools := model.Tools{}

	tools := commandsStruct.Choices(modelTools.ToolsChoice(append(toolsChoices,
		enums.Gorm, enums.Viper, enums.Wire, enums.Default)), "tools")

	//region switch case
	for _, choice := range tools {
		switch choice {
		case toolsChoices[0]:
			t.Exec(&commandsStruct, enums.Gorm, enums.GormPackage)
		case toolsChoices[1]:
			t.Exec(&commandsStruct, enums.Viper, enums.ViperPackage)
		case toolsChoices[2]:
			t.Exec(&commandsStruct, enums.Wire, enums.WirePackage)
		case enums.Default:
			fmt.Println("Saindo...")
			break
		}
	}

	//endregion
}
