package puppy

import (
	"fmt"

	"github.com/sudharsan-00/dog"
)

func From1() {
	fmt.Println("I am from version 1.0.0")
}

func From2() {
	fmt.Println("I am from version 1.2.0")
}
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
