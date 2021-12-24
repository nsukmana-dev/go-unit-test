package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Nandar")
	if result != "Hello Nandar" {
		//error
		t.Fatal("Result mush be Hello Nandar")
	}

	fmt.Println("Unit test TestHelloWorld done")
}

func TestHelloWorldSukmana(t *testing.T) {
	result := HelloWorld("Sukmana")
	if result != "Hello Sukmana" {
		//error
		t.Error("Result mush be Hello Sukmana")
	}

	fmt.Println("Unit test TestHelloWorldSukmana done")
}
