package hackerrank_test

import (
	"fmt"
	"testing"
)

type HasName interface {
	GetName() string
}

type Person struct {
	Name string
}

func SayHello(hasName HasName) string {
	return hasName.GetName()
}

func (person Person) GetName() string {
	return "Hello " + person.GetName()
}

func TestInterface(t *testing.T) {
	var melvin Person
	melvin.Name = "Melvin"

	result := SayHello(melvin)
	fmt.Println(result)
}
