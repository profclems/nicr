package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

type FileType int

const (
	Applications FileType = iota
	Archive
	Audio
	Documents
	Databases
	Fonts
	Other
	Pictures
	Videos
)

type SmartFile struct {
	Name string
	Path string
	Type FileType
}

func newRunCmd(opts *CmdOptions) *cobra.Command {
	var concurrency int
	cmd := &cobra.Command{
		Use:   "run <directory>",
		Short: "Scans and arranges files into folders according to their file types",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			dir := args[0]

			return runE(opts, dir)
		},
	}

	cmd.Flags().IntVarP(&concurrency, "concurrency", "c", 1, "NNumber of concurrency workers to use for the move operation")

	return cmd
}

func runE(opts *CmdOptions, dir string) error {
	folders := map[FileType]string{
		Applications: "Applications",
		Archive:      "Archive",
		Audio:        "Audio",
		Documents:    "Documents",
		Databases:    "Databases",
		Fonts:        "Fonts",
		Other:        "Other",
		Pictures:     "Pictures",
		Videos:       "Videos",
	}
	files, err := getFiles(dir)
	if err != nil {
		return err
	}

	if files.Len() <= 0 {
		fmt.Fprintln(opts.StdErr, "No Files found")
		return nil
	}

	//folders := make([]SmartFolder, 0, files.Len())

	for _, file := range *files {
		folder := filepath.Join(dir, folders[file.Type])
		if !dirExists(folder) {
			err := os.MkdirAll(folder, os.ModePerm)
			if err != nil {
				return err
			}
		}

		newPath := filepath.Join(folder, file.Name)
		fmt.Fprintf(opts.StdErr, "%s -> %s\n", file.Path, newPath)

		err := os.Rename(file.Path, newPath)
		if err != nil {
			return err
		}
	}
	return nil
}

type SmartFiles []*SmartFile

func (s *SmartFiles) Len() int {
	return len(*s)
}

func getFiles(dir string) (*SmartFiles, error) {
	var files SmartFiles

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
			files = append(files, &SmartFile{
				Name: v.Name(),
				Path: filepath.Join(dir, v.Name()),
				Type: getFileType(v.Name()),
			})
		}
	}

	return &files, nil
}

func getFileType(file string) FileType {
	switch filepath.Ext(file) {
	case ".jpg", ".jpeg", ".png", ".gif", ".webp", ".cr2", ".tif", ".bmp", ".heif", ".jxr", ".psd", ".ico", ".dwg":
		return Pictures

	case ".mp4", ".m4v", ".mkv", ".webm", ".mov", ".avi", ".wmv", ".mpg", ".flv", ".3gp":
		return Videos
	case ".wasm", ".dex", ".dey", ".exe", ".dmg", ".rpm", ".deb":
		return Applications
	case ".woff", ".woff2", ".ttf", ".otf":
		return Fonts
	case ".doc", ".docx", ".xls", ".xlsx", ".ppt", ".pptx", ".pdf", ".epub", ".rtf":
		return Documents
	case ".mid", ".mp3", ".m4a", ".ogg", ".flac", ".wav", ".amr", ".aac":
		return Audio
	case ".zip", ".tar", ".rar", ".gz", ".bz2", ".7z", ".xz", ".zstd", ".swf", ".iso", ".eot", ".ps", ".nes", ".crx", ".cab", ".ar", ".Z", ".lz", ".elf", ".dcm":
		return Archive
	case ".sqlite", ".sql":
		return Databases
	default:
		return Other
	}
}

func dirExists(dir string) bool {
	if info, err := os.Stat(dir); err == nil || !os.IsNotExist(err) {
		return info.IsDir()
	}
	return false
}
