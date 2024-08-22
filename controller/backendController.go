package controller

import (
	"fmt"

	"github.com/ylanzinhoy/sollievo/internal/service"
)

func BackendController() {

	var projectName string
	fmt.Println("What's your project name? ")
	fmt.Scan(&projectName)

	cf := service.CreatingFilesBackEnd{}
	cf.CompletePath = projectName
	cf.CreatingStructureBase()
	service.GomodInitService(projectName)
}
