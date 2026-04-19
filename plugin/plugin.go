package plugin

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// Info holds the metadata returned by a plugin's info command.
type Info struct {
	Name        string   `json:"name"`
	Version     string   `json:"version"`
	Engine      string   `json:"engine"`
	Description string   `json:"description"`
	Supports    []string `json:"supports"`
}

// Service defines the plugin discovery and execution interface.
type Service interface {
	Discover() ([]string, error)
	GetInfo(pluginName string) (*Info, error)
	Judge(pluginName string, filePath string, stdout io.Writer) error
}

type service struct{}

var (
	execLookPath = exec.LookPath
	execCommand  = exec.Command
	osGetenv     = os.Getenv
)

// New creates a new plugin service.
func New() Service {
	return &service{}
}

// Discover scans PATH for judge-* binaries and returns their names.
func (s *service) Discover() ([]string, error) {
	pathEnv := osGetenv("PATH")
	if pathEnv == "" {
		return nil, nil
	}

	seen := make(map[string]bool)
	var plugins []string

	for _, dir := range filepath.SplitList(pathEnv) {
		entries, err := os.ReadDir(dir)
		if err != nil {
			continue
		}
		for _, entry := range entries {
			if entry.IsDir() {
				continue
			}
			name := entry.Name()
			if strings.HasPrefix(name, "judge-") && !seen[name] {
				fullPath := filepath.Join(dir, name)
				info, err := os.Stat(fullPath)
				if err != nil {
					continue
				}
				if info.Mode()&0111 != 0 {
					seen[name] = true
					plugins = append(plugins, name)
				}
			}
		}
	}

	return plugins, nil
}

// heryEnvelope represents a HERY-wrapped JSON response.
type heryEnvelope struct {
	Type string          `json:"_type"`
	Body json.RawMessage `json:"_body"`
}

// GetInfo calls a plugin's info subcommand and parses the JSON response.
// It supports both HERY envelope format and flat JSON.
func (s *service) GetInfo(pluginName string) (*Info, error) {
	path, err := execLookPath(pluginName)
	if err != nil {
		return nil, fmt.Errorf("plugin %s not found in PATH: %w", pluginName, err)
	}

	cmd := execCommand(path, "info", "-o", "json")
	cmd.Stderr = os.Stderr
	out, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to get info from %s: %w", pluginName, err)
	}

	// Try HERY envelope first
	var envelope heryEnvelope
	if err := json.Unmarshal(out, &envelope); err == nil && envelope.Type != "" && envelope.Body != nil {
		var info Info
		if err := json.Unmarshal(envelope.Body, &info); err != nil {
			return nil, fmt.Errorf("failed to parse HERY _body from %s: %w", pluginName, err)
		}
		return &info, nil
	}

	// Fall back to flat JSON
	var info Info
	if err := json.Unmarshal(out, &info); err != nil {
		return nil, fmt.Errorf("failed to parse info from %s: %w", pluginName, err)
	}
	return &info, nil
}

// Judge calls a plugin's judge subcommand to validate data.
func (s *service) Judge(pluginName string, filePath string, stdout io.Writer) error {
	path, err := execLookPath(pluginName)
	if err != nil {
		return fmt.Errorf("plugin %s not found in PATH: %w", pluginName, err)
	}

	args := []string{"judge"}
	if filePath != "" {
		args = append(args, "-f", filePath)
	}

	cmd := execCommand(path, args...)
	cmd.Stdout = stdout
	cmd.Stderr = os.Stderr

	if filePath == "" {
		cmd.Stdin = os.Stdin
	}

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("plugin %s judge failed: %w", pluginName, err)
	}

	return nil
}
