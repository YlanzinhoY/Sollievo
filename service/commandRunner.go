package service

import (
	"fmt"
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
		log.Fatalf("Erro ao executar o comando %s\n", err)
		return err
	}

	fmt.Printf("%s instalado com sucesso! ", name)

	return nil
}
