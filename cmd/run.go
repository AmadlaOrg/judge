package cmd

import (
	"fmt"
	"os"

	"github.com/AmadlaOrg/judge/plugin"
	"github.com/spf13/cobra"
)

var (
	runFrom     string
	runFilePath string

	runPluginNew = plugin.New

	// RunCmd runs validation via a judge plugin.
	RunCmd = &cobra.Command{
		Use:   "run",
		Short: "Run validation using a judge plugin",
		Long:  "Validates data by delegating to the specified judge-* plugin (--from flag).",
		RunE:  runJudge,
	}
)

func init() {
	RunCmd.Flags().StringVar(&runFrom, "from", "", "Plugin name to run (e.g. network, application)")
	RunCmd.Flags().StringVarP(&runFilePath, "file", "f", "", "Input data file (JSON or YAML; default: stdin)")
	_ = RunCmd.MarkFlagRequired("from")
}

func runJudge(cmd *cobra.Command, args []string) error {
	pluginName := "judge-" + runFrom

	svc := runPluginNew()
	err := svc.Judge(pluginName, runFilePath, os.Stdout)
	if err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	return nil
}
