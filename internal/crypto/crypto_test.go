package crypto

import (
	"testing"
)

func TestEncode(t *testing.T) {
	var text = "Hello World!"
	var password = "password"
	var goal = "v83m5u2W0Nrp1N6b"
	var result, err = encode(text, password)
	if result != goal || err != nil {
		t.Errorf("Test Encode Failed: %s != %s", result, goal)
	}
}

func TestDecode(t *testing.T) {
	var text = "v83m5u2W0Nrp1N6b"
	var password = "password"
	var goal = "Hello World!"
	var result, err = decode(text, password)
	if result != goal || err != nil {
		t.Errorf("Test Decode Failed: %s != %s", result, goal)
	}
}
