package controller

import (
	"github.com/ylanzinhoy/sollievo/internal/model"
	"github.com/ylanzinhoy/sollievo/internal/service"
)

const (

	// names
	gorm       = "Gorm"
	viper      = "Viper"
	wire       = "Wire"
	prometheus = "Prometheus"
	fx         = "Fx"
	uuid       = "Uuid"

	// packages

	gormPackage       = "go get gorm.io/gorm"
	viperPackage      = "go get github.com/spf13/viper"
	wirePackage       = "go install github.com/google/wire/cmd/wire@latest"
	prometheusPackage = "go get github.com/prometheus/prometheus@latest"
	fxPackage         = "go get go.uber.org/fx@v1"
	uuidPackage       = "go get github.com/google/uuid"
)

func ToolingController() {

	model.NewControllerModel("tools", map[string]string{
		gorm:       gormPackage,
		viper:      viperPackage,
		wire:       wirePackage,
		prometheus: prometheusPackage,
		fx:         fxPackage,
		uuid:       uuidPackage,
	}, &service.CommandsStruct{}).ProcessCommand()
}
