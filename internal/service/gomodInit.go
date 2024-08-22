package service

import (
	"fmt"
	"log"
)

func GomodInitService(projectName string) {
	var cs *CommandsStruct
	var gomodName string

	fmt.Println("what go's mod name? ")

	_, err := fmt.Scan(&gomodName)

	if err != nil {
		log.Fatal("go mod name error")
		return
	}

	goModInit := fmt.Sprintf("mod init %s", gomodName)

	err = cs.CommandRunnerInteractivePath("go", goModInit, projectName)

	if err != nil {
		log.Fatal(err)
		return
	}

}
