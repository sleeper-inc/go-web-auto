package commands

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

var RunCommand = &cli.Command{
	Name:  "run",
	Usage: "Execute a feature file",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "feature",
			Usage:    "Path to the .feature file",
			Required: true,
		},
	},
	Action: func(c *cli.Context) error {
		featurePath := c.String("feature")
		fmt.Printf("Running feature: %s\n", featurePath)
		// Call your test runner here
		return nil
	},
}
