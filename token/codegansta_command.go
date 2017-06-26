package token

import (
	"github.com/codegangsta/cli"
	"fmt"
)

// GetTokenCommand returns the cli.Command object.
func GetTokenCommand() cli.Command {
	return cli.Command{
		Name:   "token",
		Usage:  "get oauth2 token",
		Action: CmdToken,
		Flags: 	GetTokenFlags(),
	}
}

func GetTokenFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:  fmt.Sprintf("%s, u", FLAG_USERNAME),
			Usage: "username",
		},
		cli.StringFlag{
			Name:  fmt.Sprintf("%s, p", FLAG_PASSWORD),
			Usage: "password",

		},
		cli.StringFlag{
			Name:  fmt.Sprintf("%s, i", FLAG_CLIENT_ID),
			Usage: "OAuth2 client id",
		},
		cli.StringFlag{
			Name:  fmt.Sprintf("%s, s", FLAG_CLIENT_SECRET),
			Usage: "OAuth2 client secret",
		},
	}
}
