package crypto

import (
	"crypto/sha512"
	"encoding/base64"
	"errors"
)

const (
	OFFSET byte   = 7
	HEADER string = "MESSAGE"
)

func Encode(text string, password string) (string, error) {
	if len(text) == 0 || len(password) == 0 {
		return text, errors.New("text and password must have length greater than 0")
	}
	text = HEADER + text

	h := sha512.New()
	h.Write([]byte(password))
	sha := h.Sum(nil)

	encoded := make([]byte, 0)
	for index, char := range text {
		x := byte(char)
		x += OFFSET + sha[index%len(sha)]
		encoded = append(encoded, x)
	}
	return base64.StdEncoding.EncodeToString(encoded), nil
}

func Decode(text string, password string) (string, error) {
	plain, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		return "", errors.New("invalid base64")
	}

	h := sha512.New()
	h.Write([]byte(password))
	sha := h.Sum(nil)

	decoded := make([]byte, 0)
	for index, b := range plain {
		x := b - (OFFSET + sha[index%len(sha)])
		decoded = append(decoded, x)
	}

	if len(decoded) <= len(HEADER) {
		return string(decoded), errors.New("incorrect password")
	}

	if string(decoded[0:len(HEADER)]) == HEADER {
		return string(decoded[len(HEADER):]), nil
	} else {
		return string(decoded[len(HEADER):]), errors.New("incorrect password")
	}
}
