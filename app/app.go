package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

//Returns the command line application
func Generate() *cli.App{
	app := cli.NewApp()
	app.Name = "Command Line Application"
	app.Usage = "Searches for IPs and server names on the internet"

	app.Commands = []cli.Command{
		{
			Name: "ip",
			Usage: "IPs search for internet addresses",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "host",
					Value: "google.com",
				},	
			},
			Action: searchIPs,
		},
	}

	return app
}

func searchIPs(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}