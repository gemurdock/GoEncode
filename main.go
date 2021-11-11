package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gemurdock/goencode/internal/crypto"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "GoEncode",
		Usage: "Encode and decode text with Vigenere Cipher",
		Commands: []*cli.Command{
			{
				Name:    "encrypt",
				Aliases: []string{"e"},
				Usage:   "encrypt text",
				Action: func(c *cli.Context) error {
					text, err := crypto.Encode(c.Args().First(), "test")
					if err != nil {
						return err
					}
					fmt.Println(text)
					return nil
				},
			},
			{
				Name:    "decrypt",
				Aliases: []string{"d"},
				Usage:   "decrypt text",
				Action: func(c *cli.Context) error {
					text, err := crypto.Decode(c.Args().First(), "test")
					if err != nil {
						return err
					}
					fmt.Println(text)
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
