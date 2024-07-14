package controller

import (
	"fmt"
	"net/http"

	"github.ylanzinhoy.tooling_golang/model"
	"github.ylanzinhoy.tooling_golang/service"
)

var commandsStruct = service.CommandsStruct{}
var modelTools = model.Tools{}

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
