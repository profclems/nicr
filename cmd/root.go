package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/profclems/nicr/internal/common"
)

type CmdOptions struct {
	StdErr *os.File
	StdOut *os.File
}

func NewRootCmd(opts *CmdOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   common.AppName,
		Short: "Organize files into folders by file types",
	}

	cmd.AddCommand(newRunCmd(opts))
	cmd.AddCommand(newSetupCmd(opts))
	cmd.AddCommand(newVersionCmd(opts))

	return cmd
}