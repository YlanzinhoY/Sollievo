package service

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/AlecAivazis/survey/v2"
)

func (s *CommandsStruct) WebChoice() {

	firstQuestion()

}

func firstQuestion() {

	types := []string{"react", "vue"}

	prompt := &survey.Select{
		Message: "First Choose",
		Options: types,
	}

	var frameworks string
	err := survey.AskOne(prompt, &frameworks, nil)

	cs := CommandsStruct{}

	switch frameworks {
	case "react":
		pn := createApp(&cs, "react")
		acceptTailwind(&cs, pn, "react")
	case "vue":
		pn := createApp(&cs, "vue")
		acceptTailwind(&cs, pn, "vue")
	}

	if err != nil {
		log.Fatal(err)
	}
}

func createApp(cs *CommandsStruct, frameworkName string) string {
	var pn string
	fmt.Println("Nome do projeto?")
	reader := bufio.NewReader(os.Stdin)
	pn, err := reader.ReadString('\n')

	if err != nil {
		panic(err)
	}

	pn = strings.TrimSpace(pn)

	newPn := strings.ReplaceAll(pn, " ", "_")

	command := fmt.Sprintf("pnpm create vite@latest %s  --template %s-ts", newPn, frameworkName)
	err = cs.CommandRunnerNodeJS(frameworkName, command)
	if err != nil {
		log.Fatal(err)
	}
	return newPn
}

func acceptTailwind(cs *CommandsStruct, path string, frameworkName string) {
	var choice string
	fmt.Println("Deseja Tailwind? s/n")
	_, err := fmt.Scan(&choice)

	if err != nil {
		log.Fatal(err)
	}

	if choice == strings.ToLower("s") {
		err = cs.CommandRunnerNodeJS("tailwind", fmt.Sprintf("cd %s && npm install -D tailwindcss postcss autoprefixer && npx tailwindcss init -p", path))
		if err != nil {
			log.Panic(err)
		}
		cf := &CreatingFilesBackEnd{
			path: path,
		}
		switch frameworkName {
		case "react":
			cf.TailwindReactJS()
			cf.injectTailwindAnnotationsCss("index.css")
		case "vue":
			cf.TailwindVueJS()
			cf.injectTailwindAnnotationsCss("main.css")
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
		return
	}

	if strings.ToLower(choice) == "n" {
		os.Exit(1)
	}

	backEndPath := fmt.Sprintf("%s/backend", path)

	cf := &CreatingFilesBackEnd{
		path:         path,
		completePath: backEndPath,
	}
	cf.creatingFilesBackEnd()

	var gomodName string

	fmt.Println("Nome que vc quer para seu go mod")
	_, err = fmt.Scan(&gomodName)
	if err != nil {
		log.Fatal(err)
		return
	}
	goModInit := fmt.Sprintf("mod init %s", gomodName)

	if choice == strings.ToLower("s") {
		err = cs.CommandRunnerInteractivePath("go", goModInit, backEndPath)
		if err != nil {
			fmt.Println(err)
			return
		}

		err = cs.CommandRunnerInteractivePath("sollievo", "frameworks", fmt.Sprintf("%s/backend", path))
		if err != nil {
			fmt.Println(err)
			return
		}

		err = cf.CreateFileBackend("air.toml", cf.AirToml())

		if err != nil {
			fmt.Println(err)
			return
		}

		cf.creatingStructureBase()
	}
}
