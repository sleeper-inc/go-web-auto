package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/urfave/cli/v2"
)

var ListCommand = &cli.Command{
	Name:  "list",
	Usage: "List all available .feature files in the features/ directory",
	Action: func(c *cli.Context) error {
		featureDir := "features"
		fmt.Println("Available feature files:")

		err := filepath.Walk(featureDir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() && strings.HasSuffix(info.Name(), ".feature") {
				relPath, _ := filepath.Rel(featureDir, path)
				fmt.Printf(" - %s\n", relPath)
			}
			return nil
		})

		if err != nil {
			return fmt.Errorf("failed to list features: %v", err)
		}

		return nil
	},
}
