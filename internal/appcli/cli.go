package appcli

import (
	"encoding/base64"
	"errors"
	"fmt"

	"github.com/gemurdock/goencode/internal/crypto"
	"github.com/gemurdock/goencode/internal/tools"
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
			{
				Name:    "measure",
				Aliases: []string{"m"},
				Usage:   "measure distribution of bytes for base64",
				Action: func(c *cli.Context) error {
					plain, err := base64.StdEncoding.DecodeString(c.Args().First())
					if err != nil && len(c.Args().First()) == 0 {
						return errors.New("invalid base64")
					} else if err != nil {
						plain = []byte(c.Args().First())
					}
					fmt.Println("")
					tools.PrettyDistribution(tools.GetByteDistribution(plain))
					fmt.Println("")
					return nil
				},
			},
		},
	}
}
