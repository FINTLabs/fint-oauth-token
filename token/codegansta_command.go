package token

import "github.com/codegangsta/cli"

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
			Name:  FLAG_USERNAME,
			Usage: "username",
		},
		cli.StringFlag{
			Name:  FLAG_PASSWORD,
			Usage: "password",

		},
		cli.StringFlag{
			Name:  FLAG_CLIENT_ID,
			Usage: "OAuth2 client id",
		},
		cli.StringFlag{
			Name:  FLAG_CLIENT_SECRET,
			Usage: "OAuth2 client secret",
		},
	}
}
