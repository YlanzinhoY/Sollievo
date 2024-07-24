package controller

import (
	"fmt"
)

const (
	// name

	testify = "testify"

	// package

	testifyPackage = "go get github.com/stretchr/testify"
)

func (t *ToolingControllerUpper) TestsController() {
	maps := map[string]string{
		testify: testifyPackage,
	}

	modelTools.Tools = maps

	res, err := modelTools.ToolsChoice()

	if err != nil {
		fmt.Println(err)
		return
	}

	tools := commandsStruct.Choices(res, "tests")

	t.executeChoices(maps, tools)
}
