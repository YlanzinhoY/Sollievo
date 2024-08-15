package model

import (
	"log"

	"github.com/ylanzinhoy/sollievo/internal/service"
	"github.com/ylanzinhoy/sollievo/internal/util"
)

type ControllerModel struct {
	packages       map[string]string
	commandName    string
	commandStructs *service.CommandsStruct
}

func NewControllerModel(commandName string, packages map[string]string, commandStructs *service.CommandsStruct) *ControllerModel {
	return &ControllerModel{
		commandName:    commandName,
		packages:       packages,
		commandStructs: commandStructs,
	}
}

func (c *ControllerModel) ProcessCommand() {
	res, err := c.toolsChoice()

	if err != nil {
		log.Fatal(err)
		return
	}

	tools := c.commandStructs.Choices(res, c.commandName)

	c.executeChoices(c.packages, tools)

}

func (c *ControllerModel) toolsChoice() ([]string, error) {
	var array []string
	for k := range c.packages {
		array = append(array, k)
	}

	return array, nil
}

func (c *ControllerModel) executeChoices(maps map[string]string, tools []string) {
	for _, choice := range tools {
		util.RunChoicesAndPicking(maps, choice, c.commandStructs)
	}
}
