package test

import (
	"testing"
	"new_project/hello"
	"github.com/stretchr/testify/assert"
)

// Testing using testify package (assert function)
func TestAdd2(t *testing.T) {
	expectedOutput := 3
	actualOutput := hello.Add(1, 2)
	
	assert.Equal(t, expectedOutput, actualOutput)
}

