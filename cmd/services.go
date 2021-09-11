package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
	"github.com/takama/daemon"

	"github.com/profclems/nicr/internal/common"
)

// Service is the daemon service struct
type Service struct {
	daemon.Daemon
}

type ServiceActionType int

const (
	install ServiceActionType = iota
	start
	stop
	status
	remove
)

var dependencies = []string{ /*"dummy.service"*/ }

func NewService() (*Service, error) {
	daemonKind := daemon.SystemDaemon
	if runtime.GOOS == "darwin" {
		daemonKind = daemon.UserAgent
	}
	srv, err := daemon.New(common.AppName, fmt.Sprint(common.AppName, " service"), daemonKind, dependencies...)
	if err != nil {
		return nil, err
	}

	return &Service{srv}, nil
}

func (service *Service) Manage(srvType ServiceActionType, args ...string) (string, error) {
	switch srvType {
	case install:
		return service.Install(args...)
	case remove:
		return service.Remove()
	case start:
		return service.Start()
	case stop:
		// No need to explicitly stop cron since job will be killed
		return service.Stop()
	case status:
		return service.Status()
	}

	return "", fmt.Errorf("invalid action")
}

func newServiceCmd(opts *CmdOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "service <command>",
		Short: "Manage nicr services",
	}

	installCmd := &cobra.Command{
		Use:   "install <dir>", // TODO: get dir from config
		Short: "Install the service",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			service, err := NewService()
			if err != nil {
				return err
			}

			status, err := service.Manage(install, "watch", args[0])

			if err != nil {
				return err
			}

			opts.Log.Println(status)
			return nil
		},
	}

	startCmd := &cobra.Command{
		Use:   "start",
		Short: "Start the service",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {

			service, err := NewService()
			if err != nil {
				return err
			}

			status, err := service.Manage(start)

			if err != nil {
				return err
			}

			//err = watchRun(opts, args[0])
			//if err != nil {
			//	return err
			//}

			opts.Log.Println(status)
			return nil
		},
	}

	stopCmd := &cobra.Command{
		Use:   "stop",
		Short: "Stop running services",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {

			service, err := NewService()
			if err != nil {
				return err
			}

			status, err := service.Manage(stop)

			if err != nil {
				return err
			}

			opts.Log.Println(status)
			return nil
		},
	}

	statusCmd := &cobra.Command{
		Use:   "status",
		Short: "Check the status of the service",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {

			service, err := NewService()
			if err != nil {
				return err
			}

			status, err := service.Manage(status)

			if err != nil {
				return err
			}

			opts.Log.Println(status)
			return nil
		},
	}

	removeCmd := &cobra.Command{
		Use:     "uninstall",
		Aliases: []string{"remove"},
		Short:   "Uninstall the service",
		Args:    cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {

			service, err := NewService()
			if err != nil {
				return err
			}

			status, err := service.Manage(remove)

			if err != nil {
				return err
			}

			opts.Log.Println(status)
			return nil
		},
	}

	cmd.AddCommand(installCmd, startCmd, stopCmd, statusCmd, removeCmd)

	return cmd
}
