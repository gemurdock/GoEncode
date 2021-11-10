package main

import (
	"fmt"
	"os"

	"github.com/gemurdock/goencode/internal/crypto"
)

func main() {
	var pass = "password"
	var cryptoText, err = crypto.Encode("Hello World!", pass)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(-1)
	}
	fmt.Printf("CryptoText: %s\n", cryptoText)

	var plainText string = ""
	plainText, err = crypto.Decode(cryptoText, "password")
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(-1)
	}
	fmt.Printf("PlainText: %s\n", plainText)
}
