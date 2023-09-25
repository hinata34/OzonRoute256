package top10

import (
	"bytes"
	"sort"
	"unicode"
)

func ParseWord(input []rune, index int) (string, int) {
	var result bytes.Buffer

	for ; index < len(input) && input[index] != ' ' && input[index] != '\n' && input[index] != '\t'; index++ {
		result.WriteRune(input[index])
	}

	return result.String(), index
}

func ParseWordAsterisk(input []rune, index int) (string, int) {
	var result bytes.Buffer

	for ; index < len(input) && unicode.IsLetter(input[index]) || input[index] == '-'; index++ {
		if input[index] == '-' && index > 0 && !unicode.IsLetter(input[index-1]) {
			continue
		}
		result.WriteRune(unicode.ToLower(input[index]))
	}

	return result.String(), index
}

func Top10(input string, asterisk bool) []string {
	top := make(map[string]uint64)
	givenInput := []rune(input)
	if asterisk {
		for i := 0; i < len(givenInput); {
			currWord := ""
			currWord, i = ParseWordAsterisk(givenInput, i)
			for i < len(givenInput) && !unicode.IsLetter(givenInput[i]) {
				i++
			}
			top[currWord]++
		}
	} else {
		for i := 0; i < len(givenInput); {
			currWord := ""
			currWord, i = ParseWord(givenInput, i)
			for i < len(givenInput) && (givenInput[i] == ' ' || givenInput[i] == '\n' || givenInput[i] == '\t') {
				i++
			}
			top[currWord]++
		}
	}

	type wordSet struct {
		Word   string
		Number uint64
	}
	sortedTop := []wordSet{}

	for word, number := range top {
		sortedTop = append(sortedTop, wordSet{Word: word, Number: uint64(number)})
	}

	sort.Slice(sortedTop, func(i, j int) bool {
		if sortedTop[i].Number > sortedTop[j].Number {
			return true
		} else if sortedTop[i].Number < sortedTop[j].Number {
			return false
		} else {
			return sortedTop[i].Word < sortedTop[j].Word
		}
	})

	result := []string{}
	for i := 0; i < 10 && i < len(sortedTop); i++ {
		result = append(result, sortedTop[i].Word)
	}
	return result
}
