package controller

import (
	"github.ylanzinhoy.tooling_golang/enums"
)

const (
	// name

	testify = "testify"

	// package

	testifyPackage = "go get -u github.com/stretchr/testify"
)

func (t *ToolingControllerUpper) TestsController() {
	testsChoice := make([]string, 0)

	tools := commandsStruct.Choices(modelTools.ToolsChoice(append(testsChoice, testify,
		enums.Default, enums.Back)), "tests")

	maps := map[string]string{
		testify: testifyPackage,
	}

	t.executeChoices(maps, tools)
}
