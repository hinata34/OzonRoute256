package core

import (
	"context"
	"errors"
	"homework-5/internal/app/commands"
	"os"
)

func ProcessArgs(ctx context.Context, commands map[string]commands.Command) error {
	args := os.Args
	if len(args) < 2 {
		return errors.New("input command\nfor more information use help")
	}
	command, res := commands[args[1]]

	if !res {
		return errors.New("wrong command")
	}

	err := command.Run(ctx, args[2:])

	return err
}
