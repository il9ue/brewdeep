
package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)


var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "output version information",
	Long: `
Output build version information.
`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}


var brewdeepCmd = &cobra.Command{
	Use: "brewdeep",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		context.Addr = util.EnsureHostPort(context.Addr)
	},
}

func init() {
	brewdeepCmd.AddCommand(
		initCmd,
		startCmd,
		teardownCmd,
		shutdownCmd,

		logCmd,
		statCmd,
		versionCmd,
	)
}

// Run ...
func Run(args []string) error {

}