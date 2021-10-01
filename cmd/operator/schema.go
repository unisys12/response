package operator

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/fatih/color"
	"github.com/responserms/response/internal/config"
	"github.com/responserms/response/response"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
)

var schemaFlagApply bool
var schemaFlagConfigDir string
var schemaFlagLogLevel string
var schemaLong = `
The schema command will output updates that will be made to the database on the next apply. You are expected
to take a backup before using --apply to apply the schema changes. If you do not take a backup you could lose
data if you need to revert to an older version of Response.

# Output schema changes
response operator schema

# Apply schema changes
response operator schema --apply`

var schemaApplyNotNeeded = "Your Response database does not require schema updates."
var schemaApplyCompleted = `
Response has applied the latest schema version to your database. If you need to revert to an older Response
version you should restore using the backup you created prior to applying this schema version.
`
var currentSchemaVersion = "You are currently running schema version %q. Thanks for using Response!\n"

var schema = &cobra.Command{
	Use:   "schema",
	Short: "Dump the schema or apply changes to the Response database",
	Long:  schemaLong,
	Args:  cobra.NoArgs,
	Run:   handleSchemaCommand,
}

func handleSchemaCommand(cmd *cobra.Command, args []string) {
	out := os.Stdout
	ctx, stop := signal.NotifyContext(cmd.Context(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	cfg, diags := config.NewFromGlobPath(schemaFlagConfigDir)
	if diags.HasErrors() {
		_ = diags.WriteText(out, 0, true)

		return
	}

	if schemaFlagLogLevel != "" {
		level, err := zerolog.ParseLevel(schemaFlagLogLevel)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Unable to parse the provided log level: %s", err.Error())

			return
		}

		zerolog.SetGlobalLevel(level)
	}

	app, err := response.New(cfg)
	if err != nil {
		fmt.Println(err)

		return
	}

	if err := app.Init(ctx); err != nil {
		fmt.Println(err)

		return
	}

	hasChanges, err := app.HasSchemaChanges(ctx)
	if err != nil {
		fmt.Printf("Unable to check for schema changes: %s\n", color.RedString(err.Error()))

		return
	}

	if !hasChanges {
		fmt.Println(color.BlueString(schemaApplyNotNeeded))
		fmt.Printf(color.GreenString(currentSchemaVersion), response.VersionTag)

		return
	}

	// We're only writing the changes for the schema here.
	if schemaFlagApply == false {
		if err := app.WriteSchemaChanges(ctx, nil, out); err != nil {
			fmt.Printf("Unable to write schema changes: %s\n", color.RedString(err.Error()))
		}

		return
	}

	if err := app.ApplySchemaChanges(ctx, nil); err != nil {
		fmt.Printf("Unable to apply schema changes: %s\n", color.RedString(err.Error()))
	}

	fmt.Println(color.BlueString(schemaApplyCompleted))
	fmt.Printf(color.GreenString(currentSchemaVersion), response.VersionTag)
}

func init() {
	schema.Flags().BoolVar(&schemaFlagApply, "apply", false, "Apply the necessary schema changes.")
	schema.Flags().StringVarP(&schemaFlagConfigDir, "config", "c", "config.hcl", "Glob pattern used to load and stitch configuration files")
	schema.Flags().StringVarP(&schemaFlagLogLevel, "log-level", "l", "", "Set an explicit log level to use while running this command.")
}
