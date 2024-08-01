package service

import (
	"fmt"
	"log"
	"os"
)

func creatingFilesBackEnd(path string) {
	dir := fmt.Sprintf("%s/backend", path)
	err := os.MkdirAll(dir, os.ModePerm)

	if err != nil {
		log.Fatal(err)
		return
	}

}
