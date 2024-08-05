package tests

import (
	"testing"

	"github.com/ylanzinhoy/sollievo/internal/service"
)

func TestCommandRunner(t *testing.T) {
	s := service.CommandsStruct{}

	err := s.CommandRunner("gin", "go get -u github.com/gin-gonic/gin")

	if err != nil {
		t.Error("error on instaled lib", err)
	}
}
