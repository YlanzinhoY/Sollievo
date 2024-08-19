package util

import "github.com/ylanzinhoy/sollievo/internal/service"

func RunChoicesAndPicking(libs map[string]string, value string, commandRunner *service.CommandsStruct) {
	for k, v := range libs {
		if k == value {
			err := commandRunner.CommandRunner(k, v)
			if err != nil {
				return
			}
		}
	}
}
