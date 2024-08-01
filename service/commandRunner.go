package service

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/ylanzinhoy/sollievo/enums"
)

func (s *CommandsStruct) CommandRunner(name, goModCommand string) error {

	var cmd *exec.Cmd

	cmd = exec.Command("bash", "-c", goModCommand)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	if err != nil {
		log.Fatalf(enums.Red+"Erro ao executar o comando %s\n", err)
		return err
	}

	fmt.Printf(enums.Green+"%s instalado com sucesso! ", name)

	return nil
}

func (s *CommandsStruct) CommandRunnerNodeJS(name, packageManager string) error {
	var cmd *exec.Cmd

	cmd = exec.Command("bash", "-c", packageManager)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	if err != nil {
		log.Fatalf(enums.Red+"Erro ao executar o comando %s\n", err)
		return err
	}

	fmt.Printf(enums.Green+"%s instalado com sucesso! ", name)

	return nil
}

func (s *CommandsStruct) Back(command string) error {
	var cmd *exec.Cmd

	if command == "" {
		command = "tooling_golang"
	}
	cmd = exec.Command("zsh", "-c", command)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	if err != nil {
		log.Fatalf(enums.Red+"Erro ao executar o comando %s\n", err)
		return err
	}

	return nil
}

func (s *CommandsStruct) Exit() {
	fmt.Println(enums.Purple + "Saindo...")
}
