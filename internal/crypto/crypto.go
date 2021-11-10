package crypto

import (
	"encoding/base64"
	"errors"
)

var OFFSET byte = 7

func encode(text string, password string) (string, error) {
	if len(text) == 0 || len(password) == 0 {
		return text, errors.New("text and password must have length greater than 0")
	}
	encoded := make([]byte, 0)
	for index, char := range text {
		x := byte(char)
		x += OFFSET + byte(password[index%len(password)])
		encoded = append(encoded, x)
	}
	return base64.StdEncoding.EncodeToString(encoded), nil
}

func decode(text string, password string) (string, error) {
	plain, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		return "", errors.New("invalid base64")
	}
	decoded := make([]byte, 0)
	for index, b := range plain {
		x := b - (OFFSET + byte(password[index%len(password)]))
		decoded = append(decoded, x)
	}
	return string(decoded), nil
}
