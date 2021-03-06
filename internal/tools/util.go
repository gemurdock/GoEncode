package tools

func MinMax(array []int) (int, int) {
	min := array[0]
	max := array[0]

	for _, x := range array {
		if x < min {
			min = x
		}
		if x > max {
			max = x
		}
	}
	return min, max
}

func CountOfXInt(array []int, x int) int {
	count := 0
	for _, v := range array {
		if v == x {
			count += 1
		}
	}
	return count
}

func CountOfXByte(array []byte, x byte) int {
	count := 0
	for _, v := range array {
		if v == x {
			count += 1
		}
	}
	return count
}

func NumCharLen(n int) int {
	if n < 10 && n > -10 {
		return 1
	}
	count := 0
	if n < 0 {
		n *= -1
	}
	for n > 0 {
		count++
		n /= 10
	}
	return count
}

func Total(array []int) int {
	total := 0
	for _, x := range array {
		total += x
	}
	return total
}
