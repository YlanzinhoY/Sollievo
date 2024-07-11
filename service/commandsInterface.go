package service

type CommandsInterface interface {
	CommandRunner(name, goModCommand string) error
	Choices(args []string, choice string) []string
	Back(command string) error
}

type CommandsStruct struct{}
