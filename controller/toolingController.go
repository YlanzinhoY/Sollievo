package controller

import (
	"fmt"
	"net/http"

	"github.ylanzinhoy.tooling_golang/enums"
	"github.ylanzinhoy.tooling_golang/service"
)

type ToolsServiceInterface interface {
	Exec(commandsStruct *service.CommandsStruct, value, goPackage string)
	Back(commandsStruct *service.CommandsStruct, command string)
	Exit(commandsStruct *service.CommandsStruct)
}

type ToolingControllerUpper struct{}

func (t *ToolingControllerUpper) Exec(commandsStruct *service.CommandsStruct, value, goPackage string) {
	err := commandsStruct.CommandRunner(value, goPackage)
	if err != nil {
		fmt.Println(http.StatusBadRequest)
	}
	fmt.Printf("%d\n", http.StatusOK)
}

func (t *ToolingControllerUpper) Back(commandsStruct *service.CommandsStruct, command string) {
	err := commandsStruct.Back(command)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

}

func (s *ToolingControllerUpper) Exit(commandsStruct *service.CommandsStruct) {
	fmt.Println(enums.Purple + "Saindo...")
}
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
