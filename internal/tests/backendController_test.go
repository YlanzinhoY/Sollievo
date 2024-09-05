package tests

import (
	"os"
	"testing"

	"github.com/ylanzinhoy/sollievo/internal/service"
)

func TestBackEndControllerCreatingStructureBase(t *testing.T) {
	t.Parallel()

	var services service.CreatingFilesBackEnd
	services.CompletePath = "./backend"

	if _, err := os.Stat("./backend"); os.IsNotExist(err) {

		err := os.Mkdir("./backend", 0755)
		if err != nil {
			t.Errorf("error in creating backend dir: %s", err)
			return
		}
	}
	services.CreatingStructureBase()

	defer os.RemoveAll("./backend/")
}
