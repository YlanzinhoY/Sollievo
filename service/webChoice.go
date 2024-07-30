package service

import (
	"fmt"
	"log"

	"github.com/AlecAivazis/survey/v2"
)

func (s *CommandsStruct) WebChoice() {

	questions := 2

	for questions >= 0 {
		firstQuestion()
		secondQuestion()
		questions--

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
	processAwnser(types, frameworks)
	if err != nil {
		log.Fatal(err)
	}
}

func secondQuestion() {
	types := []string{"Htmx", "react"}

	prompt := &survey.Select{
		Renderer: survey.Renderer{},
		Message:  fmt.Sprintf("Second Choose %s", types),
		Options:  types,
	}

	err := survey.AskOne(prompt, &types)

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
		}

	}

}
