package dockercomposefile

import (
	"gopkg.in/yaml.v3"
)

type (
	Config struct {
		Version  string               `yaml:"version"`
		Services map[string]Service   `yaml:"services"`
		Volumes  map[string]yaml.Node `yaml:"volumes"`
	}
	Service struct {
		Image         string            `yaml:"image"`
		ContainerName string            `yaml:"container_name"`
		Ports         []string          `yaml:"ports"`
		Environment   map[string]string `yaml:"environment"`
		Volumes       []string          `yaml:"volumes"`
	}
)

func (c *Config) Serialize(config *Config) ([]byte, error) {
	data, err := yaml.Marshal(&config)
	if err != nil {
		return nil, err
	}
	return data, nil
}
