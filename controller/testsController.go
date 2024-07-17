package controller

import (
	"fmt"

	"github.ylanzinhoy.tooling_golang/enums"
)

func (t *ToolingControllerUpper) TestsController() {
	testsChoice := make([]string, 0)

	tools := commandsStruct.Choices(modelTools.ToolsChoice(append(testsChoice, enums.Testify,
		enums.Default, enums.Back)), "tests")

	for _, choice := range tools {
		switch choice {
		case enums.Testify:
			t.Exec(&commandsStruct, enums.Testify, enums.TestifyPackage)
		case enums.Default:
			fmt.Println(enums.Purple + "Exit...")
		case enums.Back:
			t.Back(&commandsStruct, "")
		}
	}
}
