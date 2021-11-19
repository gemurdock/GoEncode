package tools

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNextPassword(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input  string
		output string
	}{
		{"a", "b"},
		{"b", "c"},
		{"c", "d"},
		{"", " "},
		{" ", "!"},
		{"/", "0"},
		{"9", ":"},
		{"@", "A"},
		{"~", "  "},
		{"aa", "ab"},
		{"~~", "   "},
		{"aR~", "aS "},
	}

	for _, test := range tests {
		assert.Equal(nextPassword(test.input), test.output, fmt.Sprintf("%s != %s", test.input, test.output))
	}
}
