package service

import (
	"fmt"
	"log"

	"github.com/AlecAivazis/survey/v2"
	"github.com/ylanzinhoy/sollievo/model"
)

func (s *CommandsStruct) WebChoice() {

	for {
		firstQuestion()

		secondQuestion()

		break

	}
}

func firstQuestion() {
	types := []string{"Htmx", "react"}

	prompt := &survey.Select{
		Renderer: survey.Renderer{},
		Message:  "First Choose",
		Options:  types,
	}

	var frameworks string
	err := survey.AskOne(prompt, &frameworks, nil)

	cs := CommandsStruct{}

	cs.CommandRunnerNodeJS("react", "npx create-react-app my-app")

	if err != nil {
		log.Fatal(err)
	}
}

func secondQuestion() {
	modeltools := model.Tools{}

	fm := model.FrameworkModel{}

	modeltools.Tools = fm.ToolsMap()

	res, err := modeltools.ToolsChoice()

	if err != nil {
		log.Panic(err)
		return
	}

	prompt := &survey.Select{
		Renderer: survey.Renderer{},
		Message:  fmt.Sprintf("qual framework voce quer? %s", res),
		Options:  res,
	}

	var choice string

	err = survey.AskOne(prompt, &choice, nil)

	if err != nil {
		log.Fatal(err)
	}

}

func processAwnser(awnsers []string, choice string) {

	length := len(awnsers) - 1

	for i := 0; i < length; i++ {
		if awnsers[i] == choice {
			// generateAwnser
			fmt.Printf("choice %s\n", choice)
			cs := CommandsStruct{}
			cs.CommandRunner(choice, "web")
		}
	}

}
