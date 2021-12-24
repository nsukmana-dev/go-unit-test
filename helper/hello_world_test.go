package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Nandar")
	if result != "Hello Nandar" {
		//error
		panic("Result is not Hello Nandar")
	}
}

func TestHelloWorldSukmana(t *testing.T) {
	result := HelloWorld("Sukmana")
	if result != "Hello Sukmana" {
		//error
		panic("Result is not Hello Sukmana")
	}
}
