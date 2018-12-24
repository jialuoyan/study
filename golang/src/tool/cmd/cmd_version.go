package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	cmdRoot.AddCommand(versionCmd)
}

var (
	ToolVer   string
	Build     string
	BuildTime string
)
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of tool",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Version: " + ToolVer)
		fmt.Println("Buid Commit: " + Build)
		fmt.Println("Buid Time: " + BuildTime)
	},
}
