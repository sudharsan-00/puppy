package puppy

import (
	"github.com/sudharsan-00/dog"
)
func Bark() string {
	return "good morning"
}

func Barks() string {
	return "had u eat lunch"
}

func Bigbark() string {
	return dog.Whengrownup(Bark())
}

func Bigbarks() string {
	return dog.Whengrownup(Barks())
}