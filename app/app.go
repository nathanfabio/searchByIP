package app

import "github.com/urfave/cli"

//Returns the command line application
func Generate() *cli.App{
	app := cli.NewApp()
	app.Name = "Command Line Application"
	app.Usage = "Searches for IPs and server names on the internet"

	return app
}