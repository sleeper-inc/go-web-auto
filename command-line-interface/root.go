package command_line_interface

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"web-auto-go-framework/command-line-interface/commands"
)

func RunCLI() {
	app := &cli.App{
		Name:  "twig",
		Usage: "Run web automation tests from the command line",
		Commands: []*cli.Command{
			commands.RunCommand,
			commands.ListCommand,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
