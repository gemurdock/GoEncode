package crypto

import (
	"testing"
)

func TestEncodeDecode(t *testing.T) {
	testText := []string{"Hello World!", "How are you doing today?", "k2YBXm}e*;,xq>A\"8&kT", "Zn/'n7uH%$8(QC2+=&AE"}
	password := "rK^\\8)rY8VQ$yP&R~4SZ"

	for _, text := range testText {
		encoded, err := Encode(text, password)
		if err != nil || len(encoded) == 0 {
			t.Errorf("Test failed: %s\n", err)
		}
		decoded, err := Decode(encoded, password)
		if err != nil || len(decoded) == 0 {
			t.Errorf("Test failed: %s\n", err)
		}
		if text != decoded {
			t.Errorf("Text does not match after encode and decode: %s -> %s\n", text, decoded)
		}
	}
}
