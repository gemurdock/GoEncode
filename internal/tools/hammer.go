package tools

import "github.com/gemurdock/goencode/internal/crypto"

type Hammer struct {
	Ciphertext string
	Passwords  []string
	Index      string
}

func nextPassword(x string) string {
	if x == "" {
		return " "
	}

	byteArray := []byte(x)

	carry := false
	for index := len(byteArray) - 1; index >= 0; index-- {
		if !carry && index < len(byteArray)-1 {
			break
		}
		if byteArray[index] == 126 {
			carry = true
			byteArray[index] = 32
		} else {
			carry = false
			byteArray[index] += 1
		}

		if index == 0 && carry {
			byteArray = append([]byte{32}, byteArray...)
		}
	}

	return string(byteArray[:])
}

func (h *Hammer) AttemptCrack() (bool, error) {
	output, err := crypto.Decode(h.Ciphertext, h.Index)

	byteOutput := []byte(output)
	isCorrect := IsASCII(byteOutput)

	if isCorrect && err == nil {
		h.Passwords = append(h.Passwords, h.Index)
	}
	h.Index = nextPassword(h.Index)

	return isCorrect, err
}
