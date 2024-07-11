package service

import (
	"fmt"
	"github.ylanzinhoy.tooling_golang/enums"
	"log"
	"os"
	"os/exec"
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

func (s *CommandsStruct) Back(command string) error {
	var cmd *exec.Cmd

	cmd = exec.Command("bash", "-c", command)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	if err != nil {
		log.Fatalf(enums.Red+"Erro ao executar o comando %s\n", err)
		return err
	}

	return nil
}
