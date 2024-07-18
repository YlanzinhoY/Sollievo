package controller

const (
	// name

	testify = "testify"

	// package

	testifyPackage = "go get -u github.com/stretchr/testify"
)

func (t *ToolingControllerUpper) TestsController() {
	maps := map[string]string{
		testify: testifyPackage,
	}

	modelTools.Tools = maps

	tools := commandsStruct.Choices(modelTools.ToolsChoice(), "tests")

	t.executeChoices(maps, tools)
}
