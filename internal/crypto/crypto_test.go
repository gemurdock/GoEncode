package crypto

import (
	"testing"
)

func TestEncode(t *testing.T) {
	var goal = "v83m5u2W0Nrp1N6b"
	var result, err = Encode("Hello World!", "password")
	if result != goal || err != nil {
		t.Errorf("Test Encode Failed: %s != %s", result, goal)
	}

	goal = "q6q6m2mefZGOSJtvu72zlpBqhoiR"
	result, err = Encode("CAPS TEXT TESTING NOW", "abcABC123!@#")
	if result != goal || err != nil {
		t.Errorf("Test Encode Failed: %s != %s", result, goal)
	}

	goal = "aseFvphopnWFvYrVo13MhrDIm7c="
	result, err = Encode("!CW^),2:<9\\u4!XKgDmW", "B}'Yh5m4")
	if result != goal || err != nil {
		t.Errorf("Test Encode Failed: %s != %s", result, goal)
	}
}

func TestDecode(t *testing.T) {
	var goal = "Hello World!"
	var result, err = Decode("v83m5u2W0Nrp1N6b", "password")
	if result != goal || err != nil {
		t.Errorf("Test Decode Failed: %s != %s", result, goal)
	}

	goal = "CAPS TEXT TESTING NOW"
	result, err = Decode("q6q6m2mefZGOSJtvu72zlpBqhoiR", "abcABC123!@#")
	if result != goal || err != nil {
		t.Errorf("Test Decode Failed: %s != %s", result, goal)
	}

	goal = "!CW^),2:<9\\u4!XKgDmW"
	result, err = Decode("aseFvphopnWFvYrVo13MhrDIm7c=", "B}'Yh5m4")
	if result != goal || err != nil {
		t.Errorf("Test Decode Failed: %s != %s", result, goal)
	}
}
