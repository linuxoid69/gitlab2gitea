package flags

import (
	"os"

	"github.com/spf13/cobra"
)

func CheckFlag(cmd *cobra.Command, flagName string) string {

	flag, err := cmd.Flags().GetString(flagName)
	if err != nil || flag == "" {
		cmd.Help()
		os.Exit(1)
	}

	return flag
}

func CheckArgs(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		cmd.Help()
	}
}
