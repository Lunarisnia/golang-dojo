package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// If no name was given return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}

	// If the name was received return the value that embeds the
	// name in the message.
	return fmt.Sprintf(randomFormat(), name), nil
}

func Hellos(names []string) (map[string]string, error) {
	if len(names) <= 0 {
		return nil, errors.New("empty names")
	}
	messages := make(map[string]string)
	for _, name := range names {
		if name == "" {
			return nil, errors.New("empty name")
		}
		messages[name] = fmt.Sprintf(randomFormat(), name)
	}
	return messages, nil
}

// init sets initial value for variables used in the function
// A lifecycle hook function that run the first time the package run
//
//	func init() {
//		rand.New(rand.NewSource(time.Now().UnixNano()))
//	}

var seed = rand.New(rand.NewSource(time.Now().UnixNano()))
var greetingsFormat = []string{
	"Hi, %v welcome.\n",
	"Great to see you, %v\n",
	"Hail, %v! Well met!\n",
	"Nice to meet ya %v!\n",
	"Hello there, %v in the flesh.\n",
}

func randomFormat() string {
	return greetingsFormat[seed.Intn(len(greetingsFormat))]
}
