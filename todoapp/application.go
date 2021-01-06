package todoapp

type Command interface {
	GetDescription() string
	GetFlag() string
}

type Application struct {
	Commands map[string]Command
}

func (a Application) Run()  {
	help := a.Commands["-h"].(*HelpCommand)
	help.Execute(a.Commands)
}
