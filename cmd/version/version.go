package version

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GitCommit = "Commit Not found"
var GitTag = "Tag Not found"

var Command = &cobra.Command{
	Use:   "version",
	Short: "a simple wildcard dns",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s: (%s)\n", GitTag, GitCommit)
	},
}
