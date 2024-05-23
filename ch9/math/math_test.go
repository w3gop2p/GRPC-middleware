package math

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFibonacci(t *testing.T) {
	actual := Fibonacci(3)
	expected := 2
	assert.Equal(t, actual, expected)
}
