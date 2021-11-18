package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculator(t *testing.T) {
	t.Run("Returns 0 when getting an empty string", func(t *testing.T) {
		actual := Add(4, 5)
		expected := 9

		assert.Equal(t, actual, expected)
	})
}
