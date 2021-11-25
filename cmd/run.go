package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"

	"github.com/profclems/nicr/internal/config"
	"github.com/profclems/nicr/internal/fileops"
)

func newRunCmd(opts *CmdOptions) *cobra.Command {
	var exclude []string
	cmd := &cobra.Command{
		Use:   "run <src> [dest]",
		Short: "Scans and arranges files into folders according to their file types",
		Args:  cobra.RangeArgs(1, 2),
		RunE: func(cmd *cobra.Command, args []string) error {
			srcDir := args[0]
			destDir := srcDir

			if len(args) == 2 {
				destDir = args[1]
			}
			return runE(opts, srcDir, destDir, exclude...)
		},
	}

	cmd.Flags().StringSliceVarP(&exclude, "exclude", "c", []string{}, "Exclude specified files or directories")

	return cmd
}

func runE(opts *CmdOptions, srcDir, destDir string, exclude ...string) error {
	filesToExclude := make(map[string]int, len(exclude))

	for i, f := range exclude {
		filesToExclude[f] = i
	}

	cfg, err := config.NewConfig(opts.ConfigPath)
	if err != nil {
		return err
	}

	files, err := getFiles(srcDir)
	if err != nil {
		return err
	}

	if files.Len() <= 0 {
		opts.Log.Println("INFO: No Files found")
		return nil
	}

	for _, file := range *files {
		if _, exclude := filesToExclude[file.Name]; exclude || file.Size <= 0 {
			continue
		}

		if folder, exclude := cfg.Get(file.Ext); !exclude {
			folder := filepath.Join(destDir, folder)
			if !fileops.DirExists(folder) {
				err := os.MkdirAll(folder, os.ModePerm)
				if err != nil {
					return err
				}
			}

			newPath := filepath.Join(folder, file.Name)
			for count := 1; ; count++ {
				newPath := newPath
				if fileops.FileExists(newPath) {
					newPath = strings.TrimRight(newPath, "."+file.Ext)
					newPath = fmt.Sprintf("%s-%d.%s", newPath, count, file.Ext)
				}
				if !fileops.FileExists(newPath) {
					opts.Log.Printf("INFO: move %s --> %s\n", file.Path, newPath)

					if err := fileops.Move(file.
						Path, newPath); err != nil {
						return err
					}
					break
				}
				continue
			}
		}
	}
	return nil
}

func getFiles(dir string) (*fileops.SmartFiles, error) {
	var files fileops.SmartFiles

	f, err := os.Open(dir)
	if err != nil {
		return nil, err
	}

	rFiles, err := f.Readdir(0)
	if err != nil {
		return nil, err
	}

	for _, v := range rFiles {
		if !v.IsDir() {
			files = append(files, &fileops.SmartFile{
				Name: v.Name(),
				Path: filepath.Join(dir, v.Name()),
				Ext:  strings.TrimLeft(filepath.Ext(v.Name()), "."),
				Size: v.Size(),
			})
		}
	}

	return &files, nil
}
