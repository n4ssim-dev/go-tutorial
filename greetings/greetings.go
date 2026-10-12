package greetings

import (
	"fmt"
	"math/rand"
	"errors"
)



func Hello(name string) (string, error) {
	//var message string
	//message = fmt.Sprint(("Hi, %v. Welcome!", name))
	if name == "" {
		return "", errors.New("Empty name.")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome !",
		"Great to see you, %v !",
		"Hail, %v, Well met !",
	}

	return formats[rand.Intn(len(formats))]
}