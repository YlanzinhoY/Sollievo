package controller

const (
	//names

	gin   = "Gin"
	echo  = "Echo"
	fiber = "Fiber"
	chi   = "Chi"

	//packages

	ginPakcage   = "go get -u github.com/gin-gonic/gin"
	fiberPackage = "go get -u github.com/gofiber/fiber/v3"
	chiPackage   = "go get -u github.com/go-chi/chi/v5"
	echoPackage  = "go get github.com/labstack/echo/v4"
)

func (t *ToolingControllerUpper) FrameworkController() {
	maps := map[string]string{
		gin:   ginPakcage,
		echo:  echoPackage,
		fiber: fiberPackage,
		chi:   chiPackage,
	}
	modelTools.Tools = maps
	tools := commandsStruct.Choices(modelTools.ToolsChoice(), "frameworks")
	t.executeChoices(maps, tools)
}
