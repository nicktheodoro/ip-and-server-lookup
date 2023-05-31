package app

import (
	"fmt"
	"net"

	"github.com/urfave/cli"
)

// GenerateApp returns a command-line application ready for execution.
func GenerateApp() *cli.App {
	app := cli.NewApp()
	app.Name = "IP and Server Lookup"
	app.Usage = "Retrieve IP addresses and server names from the internet"

	flags := []cli.Flag{
		&cli.StringFlag{
			Name:  "host",
			Value: "github.com",
			Usage: "Specify the target host",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Retrieve IP addresses from the internet",
			Flags:  flags,
			Action: retrieveIPs,
		},
		{
			Name:   "servers",
			Usage:  "Retrieve server names from the internet",
			Flags:  flags,
			Action: retrieveServers,
		},
	}

	return app
}

func retrieveIPs(c *cli.Context) error {
	host := c.String("host")

	ips, err := net.LookupIP(host)
	if err != nil {
		return err
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}

	return nil
}

func retrieveServers(c *cli.Context) error {
	host := c.String("host")

	servers, err := net.LookupNS(host)
	if err != nil {
		return err
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}

	return nil
}
