package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Nandar Sukmana Require")
	require.Equal(t, "Hello Nandar Sukmana Require", result, "Result must be 'Hello Nandar Sukmana Require'")

	fmt.Println("Unit test TestHelloWorldRequire done")
}

func TestHelloWorldAssertion(t *testing.T) {
	result := HelloWorld("Nandar Sukmana")
	assert.Equal(t, "Hello Nandar Sukmana", result, "Result must be 'Hello Nandar Sukmana'")

	fmt.Println("Unit test TestHelloWorldAssertion done")
}
