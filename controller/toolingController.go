package controller

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
	prometheusPackage = "go get github.com/prometheus/prometheus@v0.35.0"
)

func (t *ToolingControllerUpper) ToolingController() {
	toolsChoices := make([]string, 0)

	tools := commandsStruct.Choices(modelTools.ToolsChoice(append(toolsChoices,
		gorm, viper, wire, prometheus)), "tools")

	maps := map[string]string{
		gorm:       gormPackage,
		viper:      viperPackage,
		wire:       wirePackage,
		prometheus: prometheusPackage,
	}

	t.executeChoices(maps, tools)
}
