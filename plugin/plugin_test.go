package plugin

import (
	"bytes"
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestService_Discover(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "judge-test-*")
	require.NoError(t, err)
	defer os.RemoveAll(tmpDir)

	for _, name := range []string{"judge-network", "judge-application", "other-tool"} {
		f, err := os.Create(filepath.Join(tmpDir, name))
		require.NoError(t, err)
		require.NoError(t, f.Chmod(0755))
		f.Close()
	}

	origGetenv := osGetenv
	defer func() { osGetenv = origGetenv }()
	osGetenv = func(key string) string {
		if key == "PATH" {
			return tmpDir
		}
		return ""
	}

	svc := &service{}
	plugins, err := svc.Discover()
	require.NoError(t, err)

	assert.Contains(t, plugins, "judge-network")
	assert.Contains(t, plugins, "judge-application")
	assert.NotContains(t, plugins, "other-tool")
}

func TestService_DiscoverEmpty(t *testing.T) {
	origGetenv := osGetenv
	defer func() { osGetenv = origGetenv }()
	osGetenv = func(key string) string { return "" }

	svc := &service{}
	plugins, err := svc.Discover()
	require.NoError(t, err)
	assert.Nil(t, plugins)
}

func TestService_GetInfo(t *testing.T) {
	expectedInfo := Info{
		Name:        "judge-test",
		Version:     "1.0.0",
		Engine:      "test",
		Description: "Test plugin",
	}
	infoJSON, _ := json.Marshal(expectedInfo)

	origLookPath := execLookPath
	origCommand := execCommand
	defer func() {
		execLookPath = origLookPath
		execCommand = origCommand
	}()

	execLookPath = func(name string) (string, error) {
		return "/usr/bin/" + name, nil
	}
	execCommand = func(name string, args ...string) *exec.Cmd {
		return exec.Command("echo", string(infoJSON))
	}

	svc := &service{}
	info, err := svc.GetInfo("judge-test")
	require.NoError(t, err)
	assert.Equal(t, "judge-test", info.Name)
	assert.Equal(t, "test", info.Engine)
}

func TestService_GetInfoNotFound(t *testing.T) {
	origLookPath := execLookPath
	defer func() { execLookPath = origLookPath }()
	execLookPath = func(name string) (string, error) {
		return "", exec.ErrNotFound
	}

	svc := &service{}
	_, err := svc.GetInfo("judge-nonexistent")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not found in PATH")
}

func TestService_JudgeNotFound(t *testing.T) {
	origLookPath := execLookPath
	defer func() { execLookPath = origLookPath }()
	execLookPath = func(name string) (string, error) {
		return "", exec.ErrNotFound
	}

	svc := &service{}
	err := svc.Judge("judge-nonexistent", "", &bytes.Buffer{})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not found in PATH")
}
