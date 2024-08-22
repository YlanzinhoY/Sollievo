package controller

import (
	"fmt"
	"log"

	"github.com/ylanzinhoy/sollievo/internal/service"
)

func BackendController() {
	cf := service.CreatingFilesBackEnd{}
	cs := service.CommandsStruct{}

	var projectName string
	fmt.Println("What's your project name? ")
	fmt.Scan(&projectName)

	cf.CompletePath = projectName
	cf.CreatingStructureBase()
	service.GomodInitService(projectName)

	err := cs.CommandRunnerInteractivePath("sollievo", "frameworks", projectName)
	if err != nil {
		log.Fatal(err)
		return
	}
	err = cs.CommandRunnerInteractivePath("sollievo", "sqlDrivers", projectName)
	if err != nil {
		log.Fatal(err)
		return
	}
	err = cs.CommandRunnerInteractivePath("sollievo", "tools", projectName)
	if err != nil {
		log.Fatal(err)
		return
	}
	err = cs.CommandRunnerInteractivePath("sollievo", "tests", projectName)
	if err != nil {
		log.Fatal(err)
		return
	}

}
