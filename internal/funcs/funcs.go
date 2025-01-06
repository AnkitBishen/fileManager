package funcs

import (
	"log/slog"
	"os"

	"github.com/AnkitBishen/fileManagerApp/internal/types"
)

func GetFolderAndFile(path string) ([]types.DirData, error) {
	slog.Info("Get function list of files and folders", "path", path)

	dirList, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var data []types.DirData

	for _, d := range dirList {

		finfo, err := os.Stat(path + "\\" + d.Name())
		if err != nil {
			return nil, err
		}

		data = append(data, types.DirData{
			Name:         d.Name(),
			Size:         finfo.Size(),
			Type:         finfo.IsDir(),
			Path:         path + "\\" + d.Name(),
			Permission:   finfo.Mode().Perm(),
			DateModified: finfo.ModTime().Format("2006-01-02 15:04:05"),
		})

	}

	return data, nil
}
