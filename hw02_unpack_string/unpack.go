package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var res strings.Builder
	str := []rune(s)
	var buf string

	if s == "" {
		return "", nil
	}

	if !valid(s) {
		return "", ErrInvalidString
	}

	for idx, char := range str {
		currentRuneIsDigit := unicode.IsDigit(char)
		if currentRuneIsDigit {
			res.WriteString(strings.Repeat(buf, int(char-'0')))
			continue
		}
		if len(str) > idx+1 && unicode.IsDigit(str[idx+1]) {
			buf = string(char)
			continue
		}
		res.WriteRune(char)
	}
	return res.String(), nil
}

func valid(s string) bool {
	var previousRuneIsDigit bool
	str := []rune(s)

	for idx, char := range str {
		currentRuneIsDigit := unicode.IsDigit(char)
		if idx == 0 {
			if currentRuneIsDigit {
				return false
			}
			previousRuneIsDigit = currentRuneIsDigit
			continue
		}
		if currentRuneIsDigit && previousRuneIsDigit {
			return false
		}
		previousRuneIsDigit = currentRuneIsDigit
	}
	return true
}
