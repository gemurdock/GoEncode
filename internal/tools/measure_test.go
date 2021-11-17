package tools

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testPrettyDistribution struct {
	input  []byte
	output int
}

func TestPrintColor(t *testing.T) {
	assert := assert.New(t)

	resultRed, err := PrintInColorf("Hello World", "RED")
	assert.Equal(err, nil)
	assert.Equal(resultRed, 20)

	resultGreen, err := PrintInColorf("Hello World", "GREEN")
	assert.Equal(err, nil)
	assert.Equal(resultGreen, 20)

	resultWhite, err := PrintInColorf("Hello World", "WHITE")
	assert.Equal(err, nil)
	assert.Equal(resultWhite, 11)

	notIncludedColors := []string{"ORANGE", "YELLOW", "BLUE", "PURPLE", "WHITE", "BLACK", ""}
	for _, color := range notIncludedColors {
		result, err := PrintInColorf("Hello World", color)
		assert.Equal(err, nil)
		assert.Equal(result, 11)
	}

	blank, err := PrintInColorf("", "")
	assert.Equal(err, nil)
	assert.Equal(blank, 0)
}

func TestGetByteDistribution(t *testing.T) {
	assert := assert.New(t)

	byteMatrix := [][]byte{
		{10, 20, 30, 40, 50, 60, 70, 80, 90, 100},
		{10, 10, 20, 20, 30, 30, 40, 40, 50, 50},
		{0, 0, 0, 0, 0, 1, 1, 1, 1, 1},
		{},
		{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
	}

	for _, array := range byteMatrix {
		output := GetByteDistribution(array)
		for i := 0; i <= 255; i++ {
			assert.Equal(output[i], CountOfXByte(array, byte(i)))
		}
	}
}

func TestPrettyDistribution(t *testing.T) {
	assert := assert.New(t)

	tests := []testPrettyDistribution{
		{[]byte{1, 2, 3, 4, 5}, 2832},
		{[]byte{1, 1, 2, 2, 3, 3, 3, 5, 8, 25}, 2787},
		{[]byte{5, 7, 9, 22, 45, 7, 15, 7, 5, 4}, 2778},
	}

	for _, test := range tests {
		total, err := PrettyDistribution(GetByteDistribution(test.input))
		assert.Equal(err, nil)
		assert.Equal(total, test.output)
	}
}
