package controller

import (
	"github.com/ylanzinhoy/sollievo/internal/model"
	"github.com/ylanzinhoy/sollievo/internal/service"
)

const (
	// name
	testify = "testify"

	// package
	testifyPackage = "go get github.com/stretchr/testify"
)

func TestsController() {

	modelController := model.NewControllerModel("Tests", map[string]string{
		testify: testifyPackage},
		&service.CommandsStruct{},
	)
	modelController.ProcessCommand()

	//maps := map[string]string{
	//	testify: testifyPackage,
	//}
	//
	//modelTools.Tools = maps
	//
	//res, err := modelTools.ToolsChoice()
	//
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//tools := commandsStruct.Choices(res, "tests")
	//
	//t.executeChoices(maps, tools)
}
