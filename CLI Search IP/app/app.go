package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Generate returns the CLI application ready to be executed
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "CLI Application"
	app.Usage = "Search IPs and Server Names on the Internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Search IPs from internet addresses",
			Flags:  flags,
			Action: searchIps,
		},
		{
			Name:   "servers",
			Usage:  "Search the name of servers on the internet",
			Flags:  flags,
			Action: searchServers,
		},
	}

	return app
}

func searchIps(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupIP(host) // IP
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func searchServers(c *cli.Context) {
	host := c.String("host")

	servers, err := net.LookupNS(host) // Name Server
	if err != nil {
		log.Fatal(err)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}
}
