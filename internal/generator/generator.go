package generator

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/ylanzinhoy/docker_compose_gen/internal/controller"
	genfiles "github.com/ylanzinhoy/docker_compose_gen/internal/genFiles"
)

func GenerateDockerFile() {
	var types = []string{"postgres", "mysql", "mongodb"}

	prompt := &survey.Select{
		Message: "Choose a database type:",
		Options: types,
	}

	var dbType string

	survey.AskOne(prompt, &dbType)

	switch dbType {
	case "postgres":
		controller.ControllerGenerator("docker-compose.yml", genfiles.Postgres())
	case "mysql":
		controller.ControllerGenerator("docker-compose.yml", genfiles.Mysql())
	case "mongodb":
		controller.ControllerGenerator("docker-compose.yml", genfiles.MongoDB())
	}
}
