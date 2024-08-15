package dockercomposefile

func CassandraDockerFile() string {
	return `version: '3'
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

}

func PostgresDockerFile() string {
	return `version '3'
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
}