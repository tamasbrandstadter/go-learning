package todoapp

type DeleteCommand struct {
	description string
	flag        string
}

func NewDeleteCommand() Command {
	return &DeleteCommand{
		description: "Delete todo",
		flag:        "-d",
	}
}

func (d DeleteCommand) GetDescription() string {
	return d.description
}

func (d DeleteCommand) GetFlag() string {
	return d.flag
}

func (d DeleteCommand) Execute(index int64) {

}
