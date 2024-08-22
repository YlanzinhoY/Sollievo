package dockercomposefile

import (
	"fmt"
	"os"
)

type IDockerGen interface {
	generateDockerFile() error
	generateDockerFileWithPath(path string) error
}

type DockerGen struct{}

func NewDockerGen() *DockerGen {
	return &DockerGen{}
}

func (dkg *DockerGen) generateDockerFile(dataFile []byte) error {
	arquivoName := "docker-compose.yml"

	err := os.WriteFile(arquivoName, dataFile, 0666)

	if err != nil {
		return err
	}

	return nil
}

func (dkg *DockerGen) MysqlDockerFile() {
	config := Config{
		Version: "3.8",
		Services: map[string]Service{
			"mysql": {
				Image:         "mysql:latest",
				ContainerName: "mysql-container",
				Ports:         []string{"3306:3306"},
				Environment: map[string]string{
					"MYSQL_ROOT_PASSWORD": "rootpassword",
					"MYSQL_DATABASE":      "mydatabase",
					"MYSQL_USER":          "user",
					"MYSQL_PASSWORD":      "password",
				},
			},
		},
	}

	data, err := config.Serialize(&config)

	if err != nil {
		fmt.Println(err)
		return
	}

	dkg.generateDockerFile(data)
}

func (dkg *DockerGen) PostgresDockerFile() {
	config := Config{
		Version: "3.8",
		Services: map[string]Service{
			"postgres": {
				Image:         "postgres:latest",
				ContainerName: "postgres-container",
				Ports:         []string{"5432:5432"},
				Environment: map[string]string{
					"POSTGRES_USER":     "postgres",
					"POSTGRES_PASSWORD": "postgres",
					"POSTGRES_DB":       "postgres",
				},
			},
		},
	}

	data, err := config.Serialize(&config)

	if err != nil {
		fmt.Println(err)
		return
	}

	dkg.generateDockerFile(data)
}

func (dkg *DockerGen) MongoDbDockerFile() {
	config := Config{
		Version: "3.8",
		Services: map[string]Service{
			"mongodb": {
				Image:         "mongo:latest",
				ContainerName: "mongodb-container",
				Ports:         []string{"27017:27017"},
				Environment: map[string]string{
					"MONGO_INITDB_ROOT_USERNAME": "admin",
					"MONGO_INITDB_ROOT_PASSWORD": "password",
				},
			},
		},
	}

	data, err := config.Serialize(&config)

	if err != nil {
		fmt.Println(err)
		return
	}

	dkg.generateDockerFile(data)
}
