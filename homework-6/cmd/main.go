package main

import (
	"context"
	"fmt"
	commands "homework-5/internal/app/commands"
	"homework-5/internal/app/core"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	commandsCreated, err := commands.InitCommands()
	if err != nil {
		return
	}

	err = core.ProcessArgs(ctx, commandsCreated)

	if err != nil {
		fmt.Println(err)
		return
	}
}
