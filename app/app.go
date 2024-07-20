package app

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"net"
)

// Build will return the cli.Application
func Build() *cli.App {
	app := cli.NewApp()
	app.Name = "CLI Application"
	app.Usage = "Search IPs and server name on internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "devbook.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Search IPs on internet",
			Flags:  flags,
			Action: getIps,
		},
		{
			Name:   "server",
			Usage:  "Search Servers on internet",
			Flags:  flags,
			Action: getServers,
		},
	}

	return app
}

func getServers(c *cli.Context) {
	host := c.String("host")

	servers, err := net.LookupNS(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}
}

func getIps(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}
