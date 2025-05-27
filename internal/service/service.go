package service

import (
	"errors"
	//"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
	"pkg/morse"
	"strings"
)

func AutoDefinition(input string) (string, error) {
	input = strings.TrimSpace(input)
	if input == "" {
		return "", errors.New("Строка пустая")
	}
	for _, simbol := range input {
		if simbol == '.' || simbol == '-' || simbol == ' ' || simbol == '/' {

			return morse.ToText(input), nil
		}
	}
	return morse.ToMorse(input), nil
}
