package service

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"log"
	"strings"
)

func (s *CommandsStruct) WebChoice() {

	firstQuestion()

}

func firstQuestion() {

	types := []string{"Htmx", "react"}

	prompt := &survey.Select{
		Message: "First Choose",
		Options: types,
	}

	var frameworks string
	err := survey.AskOne(prompt, &frameworks, nil)

	cs := CommandsStruct{}

	switch frameworks {
	case "react":
		pn := createReactApp(&cs)
		acceptTailwind(&cs, pn)
	}

	if err != nil {
		log.Fatal(err)
	}
}

func createReactApp(cs *CommandsStruct) string {
	var pn string
	fmt.Println("Nome do projeto?")
	_, err := fmt.Scan(&pn)

	command := fmt.Sprintf("pnpm create vite@latest %s  --template react-ts", pn)
	err = cs.CommandRunnerNodeJS("react", command)
	if err != nil {
		log.Fatal(err)
	}
	return pn
}

func acceptTailwind(cs *CommandsStruct, path string) {
	var choice string
	fmt.Println("Deseja Tailwind? s/n")
	_, err := fmt.Scan(&choice)

	if err != nil {
		log.Fatal(err)
	}

	if choice == strings.ToUpper("s") || choice == strings.ToLower("s") {
		err = cs.CommandRunnerNodeJS("tailwind", fmt.Sprintf("cd %s && npm install -D tailwindcss postcss autoprefixer && npx tailwindcss init -p", path))
		if err != nil {
			log.Panic(err)
		}
	}

	acceptBackend(cs, path)
}

func acceptBackend(cs *CommandsStruct, path string) {
	var choice string
	fmt.Println("Deseja Backend? s/n")
	_, err := fmt.Scan(&choice)

	if err != nil {
		log.Fatal(err)
	}

	creatingFilesBackEnd(path)

	if choice == strings.ToLower("s") {
		err = cs.CommandRunnerInteractivePath("go", "mod init backend", fmt.Sprintf("%s/backend", path))
		if err != nil {
			fmt.Println(err)
			return
		}

		err = cs.CommandRunnerInteractivePath("sollievo", "frameworks", fmt.Sprintf("%s/backend", path))
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
