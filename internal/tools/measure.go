package tools

import (
	"fmt"
	"strconv"
	"strings"
)

func PrintInColorf(text string, color string) (int, error) {
	colorReset := "\033[0m"
	colorRed := "\033[31m"
	colorGreen := "\033[32m"

	color = strings.ToUpper(color)
	colorSwitched := true
	totalRead := 0

	switch color {
	case "RED":
		n, err := fmt.Print(colorRed)
		if err != nil {
			return n, err
		}
		totalRead += n
	case "GREEN":
		n, err := fmt.Print(colorGreen)
		if err != nil {
			return n, err
		}
		totalRead += n
	default:
		colorSwitched = false
	}

	n, err := fmt.Print(text)
	if err != nil {
		return totalRead + n, err
	}
	totalRead += n
	if colorSwitched {
		n, err := fmt.Print(colorReset)
		if err != nil {
			return totalRead + n, err
		}
		totalRead += n
	}
	return totalRead, nil
}

func GetByteDistribution(input []byte) []int {
	distribution := make([]int, 256)
	for _, x := range input {
		distribution[int(x)] += 1
	}
	return distribution
}

func GetFrequency(input []int) []float32 {
	total := Total(input)

	output := make([]float32, len(input))

	for i, x := range input {
		output[i] = float32(x) / float32(total)
	}

	return output
}

func IsASCII(input []byte) bool {
	frequency := GetFrequency(GetByteDistribution(input))

	var percentASCII float32 = 0.0

	for i, x := range frequency {
		if i >= 32 && i <= 126 {
			percentASCII += x
		}
	}

	return percentASCII > 0.95
}

func PrettyDistribution(input []int) (int, error) {
	min, max := MinMax(input)
	total := 0

	for index, x := range input {
		text := fmt.Sprintf(" %"+strconv.Itoa(NumCharLen(max))+"d", x)
		if (index+1)%16 == 0 {
			text += "\n"
		}
		if x == max {
			n, err := PrintInColorf(text, "RED")
			if err != nil {
				return n + total, err
			}
			total += n
		} else if x == min {
			n, err := PrintInColorf(text, "GREEN")
			if err != nil {
				return n + total, err
			}
			total += n
		} else {
			n, err := PrintInColorf(text, "WHITE")
			if err != nil {
				return n + total, err
			}
			total += n
		}
	}
	return total, nil
}
