package appcli

import (
	"encoding/base64"
	"errors"
	"fmt"
	"math"
	"time"

	"github.com/gemurdock/goencode/internal/crypto"
	"github.com/gemurdock/goencode/internal/tools"
	"github.com/urfave/cli/v2"
)

func StartCLI() *cli.App {
	return &cli.App{
		Name:    "GoEncode",
		Version: "alpha",
		Usage:   "Encode and decode text with Vigenere Cipher",
		Commands: []*cli.Command{
			{
				Name:    "encrypt",
				Aliases: []string{"e"},
				Usage:   "encrypt \"TEXT\" \"PASSWORD\"",
				Action: func(c *cli.Context) error {
					text, err := crypto.Encode(c.Args().Get(0), c.Args().Get(1))
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
				Usage:   "decrypt \"TEXT\" \"PASSWORD\"",
				Action: func(c *cli.Context) error {
					text, err := crypto.Decode(c.Args().Get(0), c.Args().Get(1))
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
			{
				Name:    "hammer",
				Aliases: []string{"hh"},
				Usage:   "hammer ciphertext -- Must be base64",
				Action: func(c *cli.Context) error {
					_, err := base64.StdEncoding.DecodeString(c.Args().First())
					if err != nil {
						return errors.New("invalid base64")
					}
					fmt.Printf("Cracking...\n\n")

					xh := &tools.Hammer{
						Ciphertext: c.Args().First(),
						Passwords:  []string{},
						Index:      " "}

					start := time.Now()
					lastPrint := time.Now()
					for i := int64(0); i < int64(math.Pow(256, 4)); i++ {
						isSuccess, err := xh.AttemptCrack()
						if isSuccess && err == nil {
							fmt.Printf("Possible: %s\n", xh.Passwords[len(xh.Passwords)-1])
						}
						n := time.Now()
						if n.Sub(lastPrint) > time.Second*15 {
							lastPrint = n
							totalSeconds := int64(n.Sub(start) / time.Second)
							fmt.Printf("Time: %ds @ %d hps\n", totalSeconds, i/totalSeconds)
						}
					}

					fmt.Printf("\n\nResults: %s\n", xh.Passwords)
					fmt.Printf("Len: %d\n", len(xh.Passwords))
					return nil
				},
			},
		},
	}
}
