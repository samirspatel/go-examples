package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name not allowed")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	// a map to associate names with messages
	messages := make(map[string]string)
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		//associate the retrieved name with the map
		messages[name] = message
	}
	return messages, nil
}

func randomFormat() string {
	// a slice of message formats
	formats := []string{
		"Hi %v, welcome",
		"good to see you %v",
		"kiss my ass %v",
	}
	numFormats := len(formats)
	return formats[rand.Intn(numFormats)]
}
