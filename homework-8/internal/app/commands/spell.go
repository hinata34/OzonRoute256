package commands

import (
	"context"
	"errors"
	"fmt"
)

type SpellService struct {
}

func NewSpellService() *SpellService {
	return &SpellService{}
}

func (s *SpellService) Run(ctx context.Context, args []string) error {
	if len(args) == 0 {
		return errors.New("lack of arguments")
	}

	word := args[0]
	for _, c := range word {
		if (c < 'a' || c > 'z') && (c < 'A' || c > 'Z') {
			return errors.New("not english word")
		}
	}

	for _, c := range word {
		fmt.Printf("%c ", c)
	}
	fmt.Println()
	return nil
}

func (s *SpellService) Description() string {
	return "Split english word by space"
}
