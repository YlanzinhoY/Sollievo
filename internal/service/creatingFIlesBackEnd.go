package service

import (
	"fmt"
	"log"
	"os"
)

type CreatingFilesBackEnd struct {
	path string
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

}
