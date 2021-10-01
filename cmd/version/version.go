package version

import (
	"fmt"

	"github.com/responserms/response/response"
	"github.com/spf13/cobra"
)

var versionFlagShort bool
var versionLong = `
Version returns information about the current Response version. You can use the version
command to gather important information about your version, determine if an update is
needed, or to review the Git commits that went into this release.

# Return only the version
response version --short

# Return all information
response version`

var Version = &cobra.Command{
	Use:   "version",
	Short: "Returns information about the current Response version",
	Long:  versionLong,
	Args:  cobra.NoArgs,
	Run:   handleVersionCommand,
}

func handleVersionCommand(cmd *cobra.Command, args []string) {
	if versionFlagShort {
		fmt.Println(response.VersionTag)

		return
	}

}

func init() {
	Version.Flags().BoolVar(&versionFlagShort, "short", false, "Show only the version number and exit")
}
