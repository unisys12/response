package operator

import "github.com/spf13/cobra"

var Operator = &cobra.Command{
	Use:   "operator",
	Short: "Tooling for operators of a Response deployment",
}

func init() {
	Operator.AddCommand(schema)
	Operator.AddCommand(keygen)
}
