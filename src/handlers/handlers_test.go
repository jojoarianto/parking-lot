package handlers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteLastChar(t *testing.T) {

	var tests = []struct {
		input    string
		expected string
	}{
		{"1, 2, 4, ", "1, 2, 4"},
		{"KA-01-HH-1234, KA-01-HH-9999, KA-01-P-333, ", "KA-01-HH-1234, KA-01-HH-9999, KA-01-P-333"},
		{"", ""},
	}

	assert := assert.New(t)

	for _, test := range tests {
		assert.Equal(DeleteLastChar(test.input), test.expected)
	}
}
