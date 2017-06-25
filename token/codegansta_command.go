package token

import "github.com/codegangsta/cli"

// GetTokenCommand returns the cli.Command object.
func GetTokenCommand() cli.Command {
	return cli.Command{
		Name:   "token",
		Usage:  "get oauth2 token",
		Action: CmdToken,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "username",
				Usage: "username",
			},
			cli.StringFlag{
				Name:  "password",
				Usage: "password",

			},
			cli.StringFlag{
				Name:  "client_id",
				Usage: "OAuth2 client id",
			},
			cli.StringFlag{
				Name:  "client_secret",
				Usage: "OAuth2 client secret",
			},
		},
	}
}
