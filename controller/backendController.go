package controller

import (
	"fmt"
	"log"

	"github.com/ylanzinhoy/sollievo/internal/helper"
	"github.com/ylanzinhoy/sollievo/internal/service"
)

func BackendController() {
	cf := service.CreatingFilesBackEnd{}
	cs := service.CommandsStruct{}

	var projectName string
	fmt.Println("What's your project name?")

	pn, err := helper.StdinNames(projectName)

	if err != nil {
		fmt.Println(err)
		return
	}

	cf.CompletePath = pn
	cf.CreatingStructureBase()
	service.GomodInitService(pn)

	err = cs.CommandRunnerInteractivePath("sollievo", "frameworks", pn)
	if err != nil {
		log.Fatal(err)
		return
	}
	err = cs.CommandRunnerInteractivePath("sollievo", "sqlDrivers", pn)
	if err != nil {
		log.Fatal(err)
		return
	}
	err = cs.CommandRunnerInteractivePath("sollievo", "tools", pn)
	if err != nil {
		log.Fatal(err)
		return
	}
	err = cs.CommandRunnerInteractivePath("sollievo", "tests", pn)
	if err != nil {
		log.Fatal(err)
		return
	}

}
