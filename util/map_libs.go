package util

import (
	"github.com/ylanzinhoy/sollievo/service"
)

func RunChoicesAndPicking(libs map[string]string, value string, commandRunner *service.CommandsStruct) {
	for k, v := range libs {
		if k == value {
			commandRunner.CommandRunner(k, v)
		}
	}
}
