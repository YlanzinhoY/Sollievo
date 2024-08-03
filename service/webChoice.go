package service

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	processfile "github.com/ylanzinhoy/sollievo/internal/processFile"
)

func (s *CommandsStruct) WebChoice() {

	firstQuestion()

}

func firstQuestion() {

	types := []string{"react", "flutter"}

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

	if err != nil {
		panic(err)
	}

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

	if choice == strings.ToLower("s") {
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
		return
	}

	creatingFilesBackEnd(path)

	var gomodName string

	fmt.Println("Nome que vc quer para seu go mod")
	_, err = fmt.Scan(&gomodName)
	if err != nil {
		log.Fatal(err)
		return
	}
	goModInit := fmt.Sprintf("mod init %s", gomodName)
	backEndPath := fmt.Sprintf("%s/backend", path)

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

		generateFiles(backEndPath, ".air.toml", "internal/processFile/.air.toml")
	}
}

func generateFiles(backEndPath, filename, path string) error {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Erro ao obter o diretório de trabalho:", err)
		return err
	}
	fmt.Println("Diretório de trabalho atual:", wd)

	// Atualize o caminho do arquivo de origem conforme necessário
	srcFile := filepath.Join(wd, path)
	dstFile := filepath.Join(backEndPath, filename)

	// Adicionando depuração
	absSrcFile, err := filepath.Abs(srcFile)
	if err != nil {
		fmt.Println("Erro ao obter caminho absoluto do arquivo de origem:", err)
		return err
	}
	fmt.Println("Caminho absoluto do arquivo de origem:", absSrcFile)

	absDstFile, err := filepath.Abs(dstFile)
	if err != nil {
		fmt.Println("Erro ao obter caminho absoluto do arquivo de destino:", err)
		return err
	}
	fmt.Println("Caminho absoluto do arquivo de destino:", absDstFile)

	// Verifique se o arquivo de origem existe
	if _, err := os.Stat(absSrcFile); os.IsNotExist(err) {
		fmt.Println("Arquivo de origem não existe:", absSrcFile)
		return err
	}

	err = processfile.CopyFile(absSrcFile, absDstFile)
	if err != nil {
		fmt.Println("Erro ao copiar o arquivo:", err)
		return err
	}

	fmt.Printf("%s copiado com sucesso\n", filename)

	return nil
}
