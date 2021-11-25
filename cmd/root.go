package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/profclems/nicr/internal/common"
	"github.com/profclems/nicr/internal/config"
)

type CmdOptions struct {
	StdErr *os.File
	StdOut *os.File

	ConfigPath string
	Log        *log.Logger
}

func NewRootCmd(opts *CmdOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   common.AppName,
		Short: "Organize files into folders by file types",
	}

	cmd.PersistentFlags().StringVarP(&opts.ConfigPath, "config-path", "C", config.Path(), "Path to config file")

	cmd.AddCommand(newRunCmd(opts))
	cmd.AddCommand(newSetupCmd(opts))
	cmd.AddCommand(newServiceCmd(opts))
	cmd.AddCommand(newWatchCmd(opts))
	cmd.AddCommand(newVersionCmd(opts))

	return cmd
}
