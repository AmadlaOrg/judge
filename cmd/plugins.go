package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/AmadlaOrg/judge/plugin"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var (
	pluginsNew = plugin.New

	pluginsOutputFlag string
	pluginsHeryFlag   bool

	// PluginsCmd lists all discovered judge plugins.
	PluginsCmd = &cobra.Command{
		Use:   "plugins",
		Short: "List discovered judge plugins",
		RunE:  runPlugins,
	}
)

func init() {
	PluginsCmd.Flags().StringVarP(&pluginsOutputFlag, "output", "o", "table", "Output format: table, json, yaml")
	PluginsCmd.Flags().BoolVar(&pluginsHeryFlag, "hery", false, "Wrap output in HERY envelope (_type, _body)")
}

type pluginRow struct {
	Plugin      string `json:"plugin" yaml:"plugin"`
	Engine      string `json:"engine" yaml:"engine"`
	Version     string `json:"version" yaml:"version"`
	Description string `json:"description" yaml:"description"`
}

type heryEnvelope struct {
	Type string `json:"_type" yaml:"_type"`
	Body any    `json:"_body" yaml:"_body"`
}

func runPlugins(cmd *cobra.Command, args []string) error {
	svc := pluginsNew()

	plugins, err := svc.Discover()
	if err != nil {
		return fmt.Errorf("failed to discover plugins: %w", err)
	}

	if len(plugins) == 0 {
		fmt.Fprintln(os.Stderr, "No judge plugins found in PATH.")
		return nil
	}

	var rows []pluginRow
	for _, name := range plugins {
		info, err := svc.GetInfo(name)
		if err != nil {
			rows = append(rows, pluginRow{
				Plugin:      name,
				Engine:      "?",
				Version:     "?",
				Description: fmt.Sprintf("error: %v", err),
			})
			continue
		}
		rows = append(rows, pluginRow{
			Plugin:      name,
			Engine:      info.Engine,
			Version:     info.Version,
			Description: info.Description,
		})
	}

	return renderPlugins(os.Stdout, rows, pluginsOutputFlag, pluginsHeryFlag)
}

func renderPlugins(w io.Writer, rows []pluginRow, format string, hery bool) error {
	var data any = rows
	if hery {
		data = heryEnvelope{
			Type: "amadla.org/entity/tools/plugins@v1.0.0",
			Body: rows,
		}
	}

	switch format {
	case "json":
		out, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			return err
		}
		fmt.Fprintln(w, string(out))
	case "yaml":
		out, err := yaml.Marshal(data)
		if err != nil {
			return err
		}
		fmt.Fprint(w, string(out))
	default:
		if hery {
			fmt.Fprintln(w, "_type: amadla.org/entity/tools/plugins@v1.0.0")
		}
		table := tablewriter.NewWriter(w)
		table.Header("Plugin", "Engine", "Version", "Description")
		for _, r := range rows {
			table.Append(r.Plugin, r.Engine, r.Version, r.Description)
		}
		table.Render()
	}

	return nil
}
