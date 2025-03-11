package random

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRandomString(t *testing.T) {
	tests := []struct {
		name   string
		length int
	}{
		{
			name:   "length = 5",
			length: 5,
		},
		{
			name:   "length = 1",
			length: 1,
		},
		{
			name:   "length = 10",
			length: 10,
		},
		{
			name:   "length = 100",
			length: 100,
		},
		{
			name:   "length = 20",
			length: 20,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			str1 := NewRandomString(test.length)
			str2 := NewRandomString(test.length)

			assert.Len(t, str1, test.length)
			assert.Len(t, str2, test.length)

			// Check that two generated strings are different
			// This is not an absolute guarantee that the function works correctly,
			// but this is a good heuristic for a simple random generator.
			assert.NotEqual(t, str1, str2)
		})
	}
}
