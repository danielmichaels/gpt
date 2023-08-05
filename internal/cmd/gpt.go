package cmd

import (
	"fmt"
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var (
	bins = []string{"mods", "tgpt"}
)

type GPT struct {
	Interactive bool
	PromptPath  string
}

func NewGPT() *GPT {
	return &GPT{Interactive: false}
}

func gptEntry(_ *Z.Cmd, args ...string) error {
	g := NewGPT()
	if len(args) == 0 {
		g.Interactive = true
	}
	err := g.mods(args)
	if err != nil {
		return err
	}
	return nil
}

func tgptEntry(_ *Z.Cmd, args ...string) error {
	g := NewGPT()
	if len(args) == 0 {
		g.Interactive = true
	}
	err := g.tgpt(args)
	if err != nil {
		return err
	}
	return nil
}
func (g *GPT) tgpt(args []string) error {
	tool := "tgpt"
	if g.Interactive {
		err := Z.Exec(tool, "-m")
		if err != nil {
			return err
		}
		return nil
	}
	if len(args) > 0 {
		t := []string{tool}
		t = append(t, args...)
		err := Z.Exec(t...)
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *GPT) mods(args []string) error {
	if err := setApiKey("OPENAI_API_KEY"); err != nil {
		return err
	}
	tool := []string{"mods", "--status-text", "Yeet", "-f", "-p"}
	if g.Interactive {
		g.PromptPath = filepath.Join("/tmp", fs.Isosec())
		err := Z.Exec("vim", g.PromptPath)
		if err != nil {
			return err
		}
		txt, err := os.ReadFile(g.PromptPath)
		if err != nil {
			return err
		}
		tool = append(tool, string(txt))
		err = Z.Exec(tool...)
		if err != nil {
			return err
		}
		return nil
	}
	if len(args) > 0 {
		tool = append(tool, args[0])
		err := Z.Exec(tool...)
		if err != nil {
			return err
		}
	}
	return nil
}

func binaryExists(binaryName string) bool {
	_, err := exec.LookPath(binaryName)
	if err == nil {
		return true
	}

	// Check in the calling terminal's environment variables
	pathEnv := os.Getenv("PATH")
	pathDirs := strings.Split(pathEnv, string(os.PathListSeparator))
	for _, dir := range pathDirs {
		binaryPath := filepath.Join(dir, binaryName)
		if _, err := os.Stat(binaryPath); err == nil {
			return true
		}
	}

	return false
}

func checkRequirements() error {
	for _, v := range bins {
		if !binaryExists(v) {
			Z.ExitError(fmt.Errorf("binary not found in $PATH: %q", v))
		}
	}
	return nil
}

// getConfigValue retrieves the key from YAML file, file on disk, or environment variable.
func getConfigValue(key string) (string, error) {
	if value, err := getValueFromFile(); err == nil && value != "" {
		return strings.TrimSpace(value), nil
	}
	if value := os.Getenv(strings.ToUpper(key)); value != "" {
		return value, nil
	}
	return "", fmt.Errorf("key '%s' not found in file or environment variable", key)
}

func setApiKey(key string) error {
	v, err := getConfigValue(key)
	if err != nil {
		return err
	}
	return os.Setenv("OPENAI_API_KEY", v)
}

func getValueFromFile() (string, error) {
	key := Z.Vars.Get("token")
	return key, nil
}
