package greetings

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("passed an empty name")
	}

	return fmt.Sprintf("Hi, %v! Welcome!", name), nil
}
