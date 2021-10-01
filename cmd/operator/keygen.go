package operator

import (
	"fmt"
	"os"

	"github.com/fatih/color"

	"github.com/responserms/response/internal/config"
	"github.com/spf13/cobra"
)

var keygenResult = `
# Response has generated keys for you. You should copy/paste the lines below into your Response
# configuration file.

encryption_key = "%s"

auth {
  session_secret = "%s"
}
`

//var keygenFlagWrite bool
//var keygenFlagConfig string
var keygenLong = `
The keygen command will generate a suitable encryption key to be used by Response. The command can optionally
write the key to your config file. If a glob pattern is given, the key will be written to the first file found after
sorting the files alphabetically.

When the --write flag is provided the key will not be written to STDOUT but the command will instead write the file
that the encryption key was placed in.

# Generate and write a suitable encryption key to STDOUT
response operator keygen`

var keygen = &cobra.Command{
	Use:   "keygen",
	Short: "Generate a sufficiently random encryption key",
	Long:  keygenLong,
	Args:  cobra.NoArgs,
	Run:   handleKeygenCommand,
}

func handleKeygenCommand(cmd *cobra.Command, args []string) {
	encryptionKey, err := config.GenerateEncryptionKey()
	if err != nil {
		fmt.Printf("Unable to generate a suitable encryption key: %s\n", color.RedString(err.Error()))

		return
	}

	sessionSecret, err := config.GenerateSessionSecret()
	if err != nil {
		fmt.Printf("Unable to generate a suitable encryption key: %s\n", color.RedString(err.Error()))

		return
	}

	_, _ = fmt.Fprintf(os.Stdout, keygenResult, encryptionKey, sessionSecret)
}

func init() {
	//keygen.Flags().BoolVarP(&keygenFlagWrite, "write", "w", false, "Write the encryption key to the configuration file")
}
