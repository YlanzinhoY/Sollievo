package service

type CommandsInterface interface {
	CommandRunner(name, goModCommand string) error
	CommandRunnerNodeJS(name, packageManager string) error
	CommandRunnerInteractive(name, command string) error
	CommandRunnerInteractivePath(name, command, path string) error
	Choices(args []string, choice string) []string
	WebChoice()
	Back(command string) error
	Exit()
}

type CommandsStruct struct{}
