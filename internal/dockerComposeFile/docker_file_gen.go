package dockercomposefile

import (
	"log"
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

func (dkg *DockerGen) generateDockerFile(dataFile string) error {
	arquivoName := "docker-compose.yml"

	err := os.WriteFile(arquivoName, []byte(dataFile), 0666)

	if err != nil {
		return err
	}

	return nil
}

func (dkg *DockerGen) CassandraDockerFile() {
	data := `version: '3'
services:
  cassandra:
    image: cassandra:latest
    container_name: cassandra-container
    ports:
      - "9042:9042"
    environment:
      - CASSANDRA_USER=admin
      - CASSANDRA_PASSWORD=admin
    volumes:
      - cassandra-data:/var/lib/cassandra

volumes:
  cassandra-data:`

	err := dkg.generateDockerFile(data)

	if err != nil {
		log.Fatal(err)
		return
	}
}

func (dkg *DockerGen) PostgresDockerFile() {
	data := `version: '3'
services:
  postgres:
   image: postgres:latest
   container_name: postgres-container
   ports:
	 - "5432:5432"
	environment:
		POSTGRES_USER: postgres
		POSTGRES_PASSWORD: postgres
		POSTGRES_DB: postgres
	`

	err := dkg.generateDockerFile(data)

	if err != nil {
		log.Fatal(err)
		return
	}

}
