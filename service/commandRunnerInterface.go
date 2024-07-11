package service

type CommandsInterface interface {
	CommandRunner(name, goModCommand string) error
	Choices(args []string, choice string) []string
}

type CommandsStruct struct{}
