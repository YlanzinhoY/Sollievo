package controller

import (
	"fmt"
	"net/http"

	"github.ylanzinhoy.tooling_golang/enums"
	"github.ylanzinhoy.tooling_golang/service"
	"github.ylanzinhoy.tooling_golang/util"
)

var commandsStruct = service.CommandsStruct{}

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
	commandsStruct.Exit()
}

func (s *ToolingControllerUpper) executeChoices(maps map[string]string, tools []string) {
	for _, choice := range tools {
		util.RunChoicesAndPicking(maps, choice, &service.CommandsStruct{})

		switch choice {
		case enums.Back:
			s.Back(&commandsStruct, "")
		case enums.Default:
			s.Exit(&commandsStruct)
		}
	}
}
