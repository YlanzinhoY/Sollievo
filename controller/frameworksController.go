package controller

import (
	"fmt"
)

const (
	//names

	gin        = "Gin"
	echo       = "Echo"
	fiber      = "Fiber"
	chi        = "Chi"
	gorillaMux = "GorillaMux"

	//packages

	ginPakcage        = "go get github.com/gin-gonic/gin"
	fiberPackage      = "go get github.com/gofiber/fiber/v3"
	chiPackage        = "go get github.com/go-chi/chi/v5"
	echoPackage       = "go get github.com/labstack/echo/v4"
	gorillaMuxPackage = "go get github.com/gorilla/mux"
)

func (t *ToolingControllerUpper) FrameworkController() {
	maps := map[string]string{
		gin:        ginPakcage,
		echo:       echoPackage,
		fiber:      fiberPackage,
		chi:        chiPackage,
		gorillaMux: gorillaMuxPackage,
	}
	modelTools.Tools = maps

	res, err := modelTools.ToolsChoice()

	if err != nil {
		fmt.Println(err)
		return
	}
	tools := commandsStruct.Choices(res, "frameworks")
	t.executeChoices(maps, tools)
}
