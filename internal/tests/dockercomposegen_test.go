package tests

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	dockercomposefile "github.com/ylanzinhoy/sollievo/internal/dockerComposeFile"
)

func TestDockerComposeGen(t *testing.T) {

	dc := dockercomposefile.NewDockerGen()

	dir := "./" // Replace with the path to the directory you want to search
	dc.PostgresDockerFile()

	found := false // Variable to indicate if the file was found

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && info.Name() == "docker-compose.yml" {
			t.Logf("File 'docker-compose.yml' found at: %s", path)
			found = true

			// Delete the found file
			err := os.Remove(path)
			if err != nil {
				return fmt.Errorf("error deleting the file: %w", err)
			}
			t.Log("File 'docker-compose.yml' successfully deleted!")
			return filepath.SkipDir // Stop walking the directory after finding and deleting the file
		}

		return nil
	})

	if err != nil {
		t.Errorf("Error walking the directory: %v", err)
	} else if !found {
		t.Errorf("File 'docker-compose.yml' not found in directory: %s", dir)
	}

}
