package commands

import (
	"context"
	"fmt"
)

type HelpService struct {
}

func NewHelpService() *HelpService {
	return &HelpService{}
}

func (s *HelpService) Run(ctx context.Context, args []string) error {
	for command, description := range CommandsDescription {
		fmt.Printf("%s: %s\n", command, description)
	}
	return nil
}

func (s *HelpService) Description() string {
	return "Output info about all commands"
}
