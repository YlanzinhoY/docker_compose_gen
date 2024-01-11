package generator

import (
	"log"
	"os"

	"github.com/AlecAivazis/survey/v2"
)

func GenerateDockerFile() {
	var types = []string{"postgres", "mysql"}

	prompt := &survey.Select{
		Message: "Choose a database type:",
		Options: types,
	}

	var dbType string

	survey.AskOne(prompt, &dbType)

	switch dbType {
	case "postgres":
		generatePostgresDockerFile()
	case "mysql":
		generateMySQLDockerFile()
	}
}

func generatePostgresDockerFile() {
	dockerfile := `version: '3'
services:
  postgres:
    image: postgres:latest
    environment:
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: postgres
      POSTGRES_DB: postgres
    ports:
      - "5432:5432"
`

	caminhoArquivo := "docker-compose.yml"

	err := os.WriteFile(caminhoArquivo, []byte(dockerfile), 0666)

	if err != nil {
		panic(err)
	}

	log.Printf("Arquivo Docker Compose criado em: %s\n", caminhoArquivo)

}

func generateMySQLDockerFile() {
	dockerfile := `version: '3'
services:
  mysql:
    image: mysql:latest
    environment:
      MYSQL_ROOT_PASSWORD: rootpassword
      MYSQL_DATABASE: mysqldb
      MYSQL_USER: mysqluser
      MYSQL_PASSWORD: mysqlpassword
    ports:
      - "3306:3306"
`

	caminhoArquivo := "docker-compose.yml"

	err := os.WriteFile(caminhoArquivo, []byte(dockerfile), 0666)

	if err != nil {
		panic(err)
	}

	log.Printf("Arquivo Docker Compose criado em: %s\n", caminhoArquivo)

}
