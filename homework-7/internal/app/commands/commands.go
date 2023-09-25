package commands

import "context"

var (
	CommandsDescription = make(map[string]string)
)

type Command interface {
	Run(context.Context, []string) error
	Description() string
}

func InitCommands() (map[string]Command, error) {

	commandsCreated := map[string]Command{
		"spell": NewSpellService(),
		"help":  NewHelpService(),
		"gofmt": NewGofmtService(),
		"db":    NewDBService(),
	}

	for name, command := range commandsCreated {
		CommandsDescription[name] = command.Description()
	}

	return commandsCreated, nil
}
