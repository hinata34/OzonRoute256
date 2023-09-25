package unpack

import (
	"bytes"
	"errors"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(packedString string) (string, error) {
	var result bytes.Buffer
	var prev rune = -1

	if len(packedString) == 0 {
		return "", nil
	}

	if packedString[0] >= '0' && packedString[0] <= '9' {
		return "", ErrInvalidString
	}

	screeningFlag := false

	for _, curr := range packedString {
		if curr == '\\' {
			if prev == '\\' && !screeningFlag {
				screeningFlag = true
			} else {
				if prev >= 0 && prev < '0' || prev > '9' || screeningFlag {
					result.WriteRune(prev)
				}
				screeningFlag = false
			}
		} else if curr >= '0' && curr <= '9' {
			if !screeningFlag && prev >= '0' && prev <= '9' {
				return "", ErrInvalidString
			}

			if prev == '\\' && !screeningFlag {
				screeningFlag = true
			} else {
				result.WriteString(strings.Repeat(string(prev), int(curr-'0')))
				screeningFlag = false
			}
		} else {
			if !screeningFlag && prev == '\\' {
				return "", ErrInvalidString
			}

			if prev >= 0 && prev < '0' || prev > '9' || screeningFlag {
				result.WriteRune(prev)
				screeningFlag = false
			}
		}
		prev = curr
	}

	if prev < '0' || prev > '9' || screeningFlag {
		result.WriteRune(prev)
	}

	return result.String(), nil
}
