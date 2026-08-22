package modulo

import ( 
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Name cannot be empty") 
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
} 

func randomFormat() string {

	formats := []string{
		"Hi %v, welcome",
		"Someone else will greeet you %v",
		"Hail %v full of grace",
	}
	flen := len(formats)
	rid := rand.Intn(flen)

	return formats[rid]
}
