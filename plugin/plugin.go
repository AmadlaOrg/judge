package plugin

import (
	"github.com/AmadlaOrg/judge/util/file"
	"os"
	"path/filepath"
)

type IPlugin interface {
	FindPluginsPath(pluginsPath string) (string, error)
	FindPlugins(pluginsPath string) ([]Plugins, error)
}

type SPlugin struct{}

var (
	osGetwd      = os.Getwd
	filepathAbs  = filepath.Abs
	filepathJoin = filepath.Join
	fileExists   = file.Exists
	osMkdirAll   = os.MkdirAll
	osMkdirTemp  = os.MkdirTemp
)

// FindPluginsPath
func (s *SPlugin) FindPluginsPath(pluginsPath string) (string, error) {
	var err error
	envVarPluginsPath := os.Getenv("JUDGE_PLUGINS_PATH")
	if envVarPluginsPath != "" {
		envVarPluginsPath, err = filepathAbs(envVarPluginsPath)
		if err != nil {
			return "", err
		}
	}

	return envVarPluginsPath, nil
}

// FindPlugins
func (s *SPlugin) FindPlugins(pluginsPath string) ([]Plugins, error) {
	pluginsPath, err := s.FindPluginsPath(pluginsPath)
	if err != nil {
		return nil, err
	}

	return []Plugins{}, nil
}
