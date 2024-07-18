package controller

import (
	"fmt"
)

const (

	// names

	gorm       = "Gorm"
	viper      = "Viper"
	wire       = "Wire"
	prometheus = "Prometheus"

	// packages

	gormPackage       = "go get -u gorm.io/gorm"
	viperPackage      = "go get github.com/spf13/viper"
	wirePackage       = "go install github.com/google/wire/cmd/wire@latest"
	prometheusPackage = "go get github.com/prometheus/prometheus@latest"
)

func (t *ToolingControllerUpper) ToolingController() {

	maps := map[string]string{
		gorm:       gormPackage,
		viper:      viperPackage,
		wire:       wirePackage,
		prometheus: prometheusPackage,
	}
	modelTools.Tools = maps

	res, err := modelTools.ToolsChoice()

	if err != nil {
		fmt.Println(err)
		return
	}

	tools := commandsStruct.Choices(res, "tools")
	t.executeChoices(maps, tools)
}
