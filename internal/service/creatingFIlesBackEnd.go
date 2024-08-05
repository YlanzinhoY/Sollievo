package service

import (
	"fmt"
	"log"
	"os"
)

type CreatingFilesBackEnd struct {
	path         string
	completePath string
}

func (c *CreatingFilesBackEnd) creatingFilesBackEnd() {

	dir := fmt.Sprintf("%s/backend", c.path)
	err := os.MkdirAll(dir, os.ModePerm)

	if err != nil {
		log.Fatal(err)
		return
	}

}

func (c *CreatingFilesBackEnd) creatingStructureBase() {
	projectDir := []string{
		"cmd",
		"pkg",
		"internal",
		"api",
		"configs",
		"data",
		"scripts",
		"docs",
		"test",
	}

	for _, dir := range projectDir {
		err := os.MkdirAll(fmt.Sprintf("%s/%s", c.completePath, dir), os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("estrutura criada com sucesso!")
	}

}
