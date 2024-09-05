package helper

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func StdinNames(name string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	name = strings.TrimSpace(name)

	newName := strings.ReplaceAll(name, " ", "_")

	return newName, nil

}
