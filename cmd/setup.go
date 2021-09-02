package cmd

import (
	"github.com/spf13/cobra"
)

func newSetupCmd(opts *CmdOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "setup",
		Short: "Create nicr config file",
		Args:  cobra.MaximumNArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	return cmd
}
