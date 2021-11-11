package appcli

import (
	"fmt"

	"github.com/gemurdock/goencode/internal/crypto"
	"github.com/urfave/cli/v2"
)

func StartCLI() *cli.App {
	return &cli.App{
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
}
