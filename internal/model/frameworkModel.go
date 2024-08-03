package model

type FrameworkModel struct {
}

func (f *FrameworkModel) ToolsMap() map[string]string {
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

	maps := map[string]string{
		gorm:       gormPackage,
		viper:      viperPackage,
		wire:       wirePackage,
		prometheus: prometheusPackage,
		fx:         fxPackage,
		uuid:       uuidPackage,
	}

	return maps
}
