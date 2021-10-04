
   
package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"net"
	"os"
)

func main(){
	app := cli.NewApp()
	app.Name = "Website Lookup CLI"
	app.Usage = "Let's you query IPs, CNAMs, MX records and Name Servers"

	myFlags := []cli.Flag{
		&cli.StringFlag{
			Name: "host",
			Value: "google.com",
		},
	}

	app.Commands = []*cli.Command{
		{
			Name: "ns",
			Usage: "Looks up the Name Servers for a Particular Host",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				ns, err := net.LookupNS(c.String("host"))
				if err != nil {
					fmt.Println(err)
					return err
				}
				for i:=0; i<len(ns); i++{
					fmt.Println(ns[i].Host)
				}
				return nil
			},
		},
		{
			Name:"ip",
			Usage: "Looks Up the IP addresses for a particular host",
			Flags:myFlags,
			Action: func(c *cli.Context) error {
				ip, err := net.LookupIP(c.String("host"))
				if err !=nil {
					fmt.Println(err)
					return err
				}
				for i:=0; i<len(ip); i++ {
					fmt.Println(ip[i])

				}
				return nil
			},

		}
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)

	}
}
