package cmd

import (
    // "fmt"
	"github.com/spf13/cobra"
	"tool/shell"
)

var cmdRemove = &cobra.Command{
	Use:   "remove",
	Short: "docker container stop or remove",
	Run: func(cmd *cobra.Command, args []string) {
		shell.Execute("docker-compose -f docker-compose.yml  down --remove-orphans")
	},
}

func init() {
	cmdRoot.AddCommand(cmdRemove)
}