package tools

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testContainerCountByXInt struct {
	input  []int
	filter int
	output int
}

type testContainerCountByXByte struct {
	input  []byte
	filter byte
	output int
}

func TestMinMax(t *testing.T) {
	min, max := MinMax([]int{0, 1, 2, 3, 4, 5})
	if min != 0 {
		t.Errorf("Test failed: Wrong minimum %d should be 0\n", min)
	}
	if max != 5 {
		t.Errorf("Test failed: Wrong maximum %d should be 5\n", min)
	}

	min, max = MinMax([]int{5, 4, 3, 2, 1, 0})
	if min != 0 {
		t.Errorf("Test failed: Wrong minimum %d should be 0\n", min)
	}
	if max != 5 {
		t.Errorf("Test failed: Wrong maximum %d should be 5\n", min)
	}

	min, max = MinMax([]int{3, 2, 5, 0, 1, 4})
	if min != 0 {
		t.Errorf("Test failed: Wrong minimum %d should be 0\n", min)
	}
	if max != 5 {
		t.Errorf("Test failed: Wrong maximum %d should be 5\n", min)
	}
}

func TestCountOfXInt(t *testing.T) {
	assert := assert.New(t)

	tests := []testContainerCountByXInt{
		{[]int{1, 2, 3, 2, 4, 5}, 2, 2},
		{[]int{1, 2, 3, 2, 4, 5}, 1, 1},
		{[]int{1, 2, 3, 2, 4, 5}, 5, 1},
		{[]int{1, 2, 3, 2, 4, 5}, 0, 0},
		{[]int{1, 2, 3, 2, 4, 5}, -1, 0},
		{[]int{-1, 2, -1, 2, -1, 5}, -1, 3},
	}

	for _, test := range tests {
		assert.Equal(CountOfXInt(test.input, test.filter), test.output)
	}
}

func TestCountOfX(t *testing.T) {
	assert := assert.New(t)

	tests := []testContainerCountByXByte{
		{[]byte{1, 2, 3, 2, 4, 5}, 2, 2},
		{[]byte{1, 2, 3, 2, 4, 5}, 1, 1},
		{[]byte{1, 2, 3, 2, 4, 5}, 5, 1},
		{[]byte{1, 2, 3, 2, 4, 5}, 0, 0},
		{[]byte{1, 2, 3, 2, 4, 5}, 255, 0},
	}

	for _, test := range tests {
		assert.Equal(CountOfXByte(test.input, test.filter), test.output)
	}
}

func TestNumCharLen(t *testing.T) {
	assert := assert.New(t)

	tests := [][]int{
		{0, 1},
		{1, 1},
		{2, 1},
		{10, 2},
		{15, 2},
		{123, 3},
		{1234, 4},
		{12345, 5},
		{123456, 6},
		{99999, 5},
		{11111, 5},
		{10000, 5},
		{-0, 1},
		{-1, 1},
		{-2, 1},
		{-10, 2},
		{-15, 2},
		{-123, 3},
		{-1234, 4},
		{-12345, 5},
		{-123456, 6},
		{-99999, 5},
		{-11111, 5},
		{-10000, 5},
	}

	for _, test := range tests {
		assert.Equal(NumCharLen(test[0]), test[1])
	}
}

func TestTotal(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input  []int
		output int
	}{
		{[]int{1, 2, 3}, 6},
		{[]int{9, 5, 4, 9, 3, 2}, 32},
		{[]int{4, 8, 4, 99, -4}, 111},
		{[]int{2, 4, 6, 8}, 20},
		{[]int{0, 0, 1, 2, -5}, -2},
	}

	for _, test := range tests {
		assert.Equal(Total(test.input), test.output)
	}
}
