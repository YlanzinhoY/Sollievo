package controller

import "github.com/ylanzinhoy/sollievo/internal/service"

func BackendController() {

	


	cf := service.CreatingFilesBackEnd{}
	cf.CompletePath = "backend"

	cf.CreatingStructureBase()

}
