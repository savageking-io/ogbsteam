package main

import (
	ogb "github.com/savageking-io/ogbcommon"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "ogbsteam"
	app.Version = AppVersion
	app.Description = "Smart backend service for smart game developers"
	app.Usage = "Steam Microservice for OnlineGameBase"

	app.Authors = []cli.Author{
		{
			Name:  "savageking.io",
			Email: "i@savageking.io",
		},
		{
			Name:  "Mike Savochkin (crioto)",
			Email: "mike@crioto.com",
		},
	}

	app.Copyright = "2025 (c) savageking.io. All Rights Reserved"

	app.Commands = []cli.Command{
		{
			Name:  "serve",
			Usage: "Start database service",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "config",
					Usage:       "Configuration filepath",
					Value:       ConfigFilepath,
					Destination: &ConfigFilepath,
				},
				cli.StringFlag{
					Name:        "log",
					Usage:       "Specify logging level",
					Value:       LogLevel,
					Destination: &LogLevel,
				},
			},
			Action: Serve,
		},
	}

	_ = app.Run(os.Args)
}

func Serve(c *cli.Context) error {
	err := ogb.SetLogLevel(LogLevel)
	if err != nil {
		return err
	}

	err = ogb.ReadYAMLConfig(ConfigFilepath, &AppConfig)
	if err != nil {
		return err
	}

	log.Infof("Configuration loaded from %s", ConfigFilepath)
	return nil
}
