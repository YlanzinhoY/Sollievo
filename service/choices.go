package service

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"log"
)

func (s *CommandsStruct) Choices(args []string, choice string) []string {

	var choices = args
	var choiceType []string

	prompt := &survey.MultiSelect{
		Renderer: survey.Renderer{},
		Message:  fmt.Sprintf("Choose %s", choice),
		Options:  choices,
	}

	err := survey.AskOne(prompt, &choiceType)

	if err != nil {
		log.Fatal(err)
		return nil
	}
	return choiceType
}
