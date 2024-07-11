package controller

import (
	"fmt"
	"github.ylanzinhoy.tooling_golang/enums"
)

func (t *ToolingControllerUpper) FrameworkController() {
	frameworkChoices := make([]string, 0)

	tools := commandsStruct.Choices(modelTools.ToolsChoice(append(frameworkChoices, enums.Gin,
		enums.Echo, enums.Chi, enums.Fiber, enums.Default, enums.Back)), "frameworks")

	for _, choice := range tools {
		switch choice {
		case enums.Gin:
			t.Exec(&commandsStruct, enums.Gin, enums.GinPakcage)
		case enums.Echo:
			t.Exec(&commandsStruct, enums.Echo, enums.EchoPackage)
		case enums.Chi:
			t.Exec(&commandsStruct, enums.Chi, enums.EchoPackage)
		case enums.Fiber:
			t.Exec(&commandsStruct, enums.Fiber, enums.FiberPackage)
		case enums.Default:
			fmt.Println(enums.Purple + "Saindo...")
			break
		case enums.Back:
			t.Back(&commandsStruct, "tooling_golang")
			break
		default:
			break
		}
	}
}
