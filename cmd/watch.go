package cmd

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
)

func newWatchCmd(opts *CmdOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "watch <dir>",
		Short: "Watches the directory for changes",
		Long:  "Watches a directory for [file] changes and automatically organizes the files in the directory",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			dir := args[0]

			return watchRun(opts, dir)
		},
	}

	return cmd
}

func watchRun(opts *CmdOptions, dir string) error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				opts.Log.Println("EVENT:", event)
				if event.Op == fsnotify.Create {
					err = runE(opts, dir, dir, nil)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				opts.Log.Println("ERROR:", err)
			}
		}
	}()

	err = watcher.Add(dir)
	if err != nil {
		return err
	}
	<-done
	return nil
}
