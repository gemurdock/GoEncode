package appcli

import (
	"testing"
)

func TestEncodeDecode(t *testing.T) {
	argTests := [][]string{
		{"cipher", "encrypt", "Hello World!"},
		{"cipher", "decrypt", "w9Hm5+qM0ert2N6cc"},
	}
	for _, args := range argTests {
		app := StartCLI()
		err := app.Run(args)
		if err != nil {
			t.Errorf("Test failed: %s\n", err)
		}
	}
}

func TestDecodeErrorOnInvalidBase64(t *testing.T) {
	app := StartCLI()
	err := app.Run([]string{"cipher", "decrypt", "w9sHm5+qM0ertt2N6cc"})
	if err == nil {
		t.Errorf("Test failed: invalid base64 is not rejected\n")
	}
}
