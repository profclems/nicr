package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/mitchellh/go-homedir"

	"github.com/profclems/nicr/internal/fileops"
)

const (
	defaultUnknownExtFolderName = "Other"
)

var (
	cfg       *Config
	pathCache string
)

type Config struct {
	UnknownFilesFolder string    `json:"unknown_files_folder"`
	KnownFiles         []FileExt `json:"known_files"`

	Path string

	exemptedExts map[string]string
	includedExts map[string]string
}

type FileExt struct {
	Extensions  []string `json:"extensions"`
	Folder      string   `json:"folder"`
	ExemptFiles bool     `json:"exempt_files"`
}

// Path returns config file path
func Path() string {
	usrConfigHome := os.Getenv("XDG_CONFIG_HOME")
	if usrConfigHome == "" {
		usrConfigHome = os.Getenv("HOME")
		if usrConfigHome == "" {
			usrConfigHome, _ = homedir.Expand("~/.config")
		} else {
			usrConfigHome = filepath.Join(usrConfigHome, ".config")
		}
	}
	return filepath.Join(usrConfigHome, "nicr", "config.json")
}

func NewConfig(path string) (*Config, error) {
	if path == "" {
		path = Path()
	}

	if cfg != nil && path == pathCache {
		return cfg, nil
	}

	if !fileops.FileExists(path) {
		err := os.MkdirAll(filepath.Dir(path), 0750)
		if err != nil {
			return nil, err
		}
		err = ioutil.WriteFile(path, []byte(defaultConfig), 0600)
		if err != nil {
			return nil, fmt.Errorf("could not create config file: %w", err)
		}
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	cfg = &Config{
		UnknownFilesFolder: defaultUnknownExtFolderName,
	}

	err = json.Unmarshal(data, cfg)
	if err != nil {
		return nil, err
	}

	pathCache = path
	cfg.Path = path

	return cfg, nil
}

func (c *Config) initCfgMap() {
	c.includedExts = make(map[string]string)
	c.exemptedExts = make(map[string]string)

	for _, file := range c.KnownFiles {
		for _, extension := range file.Extensions {
			extension = strings.TrimLeft(extension, ".")
			if file.ExemptFiles {
				c.exemptedExts[extension] = file.Folder
				continue
			}
			c.includedExts[extension] = file.Folder
		}
	}
}

// Get returns the folder as string and a boolean val indicating whether the
// ext should be excluded or not.
func (c *Config) Get(extension string) (string, bool) {
	if len(c.includedExts) == 0 && len(c.exemptedExts) == 0 {
		c.initCfgMap()
	}

	if folder, ok := c.exemptedExts[extension]; ok {
		return folder, true
	}

	if folder, ok := c.includedExts[extension]; ok {
		return folder, false
	}

	return c.UnknownFilesFolder, false
}
