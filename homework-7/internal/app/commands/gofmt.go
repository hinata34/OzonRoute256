package commands

import (
	"bufio"
	"context"
	"errors"
	"os"
	"strings"
	"unicode"
)

type GofmtService struct {
}

func NewGofmtService() *GofmtService {
	return &GofmtService{}
}

func (s *GofmtService) Run(ctx context.Context, args []string) error {
	if len(args) == 0 {
		return errors.New("lack of arguments")
	}

	fileName := args[0]
	inFile, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer inFile.Close()

	outFile, err := os.Create(fileName + ".formated")
	if err != nil {
		return err
	}
	defer outFile.Close()

	reader := bufio.NewScanner(inFile)
	reader.Split(bufio.ScanLines)

	var line string
	var words []string
	if reader.Scan() {
		line = reader.Text()
		words = strings.Split(line, " ")
		var parsedWords []string

		for _, word := range words {
			if word != "" {
				parsedWords = append(parsedWords, word)
			}
		}
		words = parsedWords
	} else {
		return nil
	}

	for {
		for i := 1; i < len(words); i++ {
			if unicode.IsUpper(rune(words[i][0])) {
				outFile.WriteString(words[i-1] + ". ")
			} else {
				outFile.WriteString(words[i-1] + " ")
			}
		}
		if !reader.Scan() {
			outFile.WriteString(words[len(words)-1] + "\n")
			break
		}
		nextLine := reader.Text()
		nextWords := strings.Split(nextLine, " ")

		var parsedNextWords []string

		for _, word := range nextWords {
			if word != "" {
				parsedNextWords = append(parsedNextWords, word)
			}
		}

		if len(parsedNextWords) > 0 && unicode.IsUpper(rune(parsedNextWords[0][0])) {
			outFile.WriteString(words[len(words)-1] + ".\n")
		} else {
			outFile.WriteString(words[len(words)-1] + "\n")
		}

		line = nextLine
		words = parsedNextWords
	}
	return nil
}

func (s *GofmtService) Description() string {
	return "Formats the file"
}
