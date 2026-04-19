package main

import (
	"fmt"
	"os"

	"github.com/AmadlaOrg/judge/cmd"
	"github.com/spf13/cobra"
)

const (
	appName = "judge"
	version = "1.0.0"
)

var rootCmd = &cobra.Command{
	Use:     appName,
	Short:   "Validation and audit CLI with judge-* plugins",
	Version: version,
}

func init() {
	rootCmd.AddCommand(cmd.RunCmd)
	rootCmd.AddCommand(cmd.PluginsCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
