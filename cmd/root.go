package cmd

import (
	"github.com/responserms/response/cmd/operator"
	"github.com/responserms/response/cmd/serve"
	"github.com/responserms/response/cmd/version"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "response",
	Short: "Open source CAD, MDT, and RMS.",
	Long: `Response is a free and open source CAD, MDT, and RMS product designed
to be developer-friendly, heavily customizable, and fast.

# Start Response
response response serve

# Start Response in Development Mode
response response serve dev

# Update Response
response update

# Check the current Response version
response version
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	//  Serve
	rootCmd.AddCommand(serve.Serve)

	//  Update
	//rootCmd.AddCommand(update.Update)

	//  Version
	rootCmd.AddCommand(version.Version)

	// Operator
	rootCmd.AddCommand(operator.Operator)
}
