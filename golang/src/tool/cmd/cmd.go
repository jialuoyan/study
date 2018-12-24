package cmd
import (
  "fmt"
  "os"
 
  "github.com/spf13/cobra"
)
 
var (
	rootFlagLatest  bool
	rootFlagVerbose bool
	rootFlagPrefix  string
)

var cmdRoot = &cobra.Command{
	Use:   "tool",
	Short: "this is a tool",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
	    fmt.Println("all command will print this string")
	},
}

func Execute() {
	if err := cmdRoot.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}


func init() {
	cmdRoot.PersistentFlags().BoolVarP(&rootFlagLatest, "latest", "", false, "version to latest")
	cmdRoot.PersistentFlags().StringVarP(&rootFlagPrefix, "prefix", "", "", "prefix -")
	cmdRoot.PersistentFlags().BoolVarP(&rootFlagVerbose, "verbose", "v", false, "print detail infomaton")
}
