package serve

import (
	"context"
	"fmt"
	"os"

	"github.com/rs/zerolog"

	"github.com/fatih/color"
	"github.com/responserms/response/server"

	"github.com/contaim/spec"
	_ "github.com/fatih/color"
	"github.com/responserms/response/internal/config"
	"github.com/responserms/response/response"
	"github.com/spf13/cobra"
)

var serveFlagPprof bool
var serveFlagConfigDir string
var serveFlagLogLevel string
var serveFlagApplySchema bool
var serveLong = `
The serve command (and sub-commands) start the Response server, which includes the GraphQL API, Response Console, and
other backend APIs as one single process.

# Start Response with a config.hcl file in the current directory
response serve

# Start Response with a response.hcl file in the current directory
response serve --config=response.hcl

# Start Response and merge all HCL config files in the ./config directory
response serve --config=config/*.hcl

# Start Response with debug logging
response serve --log-level=debug

# Start Response with profiling endpoints enabled
response serve --pprof
`
var Serve = &cobra.Command{
	Use:   "serve",
	Short: "Serve starts and manages a Response instance.",
	Long:  serveLong,
	Args:  cobra.MaximumNArgs(1),
	Run:   handleServe,
}

func handleServe(cmd *cobra.Command, args []string) {
	var diag *spec.Diagnostics
	var cfg *config.Config
	var err error

	ctx, cancel := context.WithCancel(cmd.Context())
	defer cancel()

	cfg, diag = config.NewFromGlobPath(serveFlagConfigDir)
	if diag.HasErrors() {
		_ = diag.WriteText(os.Stdout, 0, true)

		return
	}

	if serveFlagLogLevel != "" {
		level, err := zerolog.ParseLevel(serveFlagLogLevel)
		if err != nil {
			fmt.Printf("Invalid log level provided: %s", err.Error())

			return
		}

		cfg.LogLevel = level
	}

	if serveFlagPprof {
		cfg.DeveloperProfiling = true
	}

	app, err := response.New(cfg)
	if err != nil {
		fmt.Printf("There was a problem while initializing the Response app: %s\n", color.RedString(err.Error()))

		return
	}

	if err := app.Init(ctx); err != nil {
		if err != response.ErrSchemaChangesAvailable {
			fmt.Printf("Response failed to initialize: %s\n", color.RedString(err.Error()))

			return
		}

		if err == response.ErrSchemaChangesAvailable && !serveFlagApplySchema {
			fmt.Printf("Response cannot start: %s\n", color.RedString("There are unapplied schema changes in your database. Apply the latest schema and then start Response."))

			return
		}
	}

	if serveFlagApplySchema {
		_, _ = fmt.Fprintf(os.Stdout, "%s\n", color.RedString("!!! Response is applying the latest schema before starting. Do NOT do this in production, you should take backups!"))

		if err := app.ApplySchemaChanges(ctx, &response.SchemaChangeConfig{}); err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "\n\nResponse failed to apply the latest schema: %s\n", color.RedString(err.Error()))

			return
		}
	}

	if err != nil {
		fmt.Printf("response.New: %s", err.Error())
	}

	if !cfg.HideBanner {
		WriteBanner(cmd.OutOrStdout(), response.VersionTag)
	}

	svr, err := server.New(app)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "\nResponse HTTP Server has exited: %s\n", color.RedString(err.Error()))

		return
	}

	app.RegisterMonitoredService(svr.MonitoredService(ctx))

	if err := app.Run(ctx); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "\n\nResponse has exited: %s\n", color.RedString(err.Error()))
	}
}

func init() {
	Serve.Flags().StringVarP(&serveFlagConfigDir, "config", "c", "config.hcl", "Glob pattern used to load and stitch configuration files")
	Serve.Flags().StringVarP(&serveFlagLogLevel, "log-level", "l", "", "Start Response with a specific log level. Overrides configuration file.")
	Serve.Flags().BoolVarP(&serveFlagApplySchema, "apply-schema", "a", false, "Apply the schema changes before starting")
	Serve.Flags().BoolVar(&serveFlagPprof, "pprof", false, "Enable pprof endpoints at /debug/pprof")
}
