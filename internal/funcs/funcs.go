package funcs

import (
	"log/slog"
	"os"
	"strings"

	"github.com/AnkitBishen/fileManagerApp/internal/types"
)

type Result struct {
	TotalFiles   int             `json:"totalFiles"`
	TotalFolders int             `json:"totalFolders"`
	DirData      []types.DirData `json:"dirData"`
}

// GetFolderAndFile get list of files and folders
// params path string and function return list of files and folders or error
func GetFolderAndFile(path string) (Result, error) {
	slog.Info("Get function list of files and folders", "path", path)

	dirList, err := os.ReadDir(path)
	if err != nil {
		return Result{}, err
	}

	var data []types.DirData
	var totalFiles int = 0
	var totalFolders int = 0

	for _, d := range dirList {

		finfo, err := os.Stat(path + "\\" + d.Name())
		if err != nil {
			return Result{}, err
		}

		// convert file size in bytes to KB
		fSizeInBytes := finfo.Size()
		fSizeInKB := int64(fSizeInBytes) / 1024

		// fileMode := fileInfo.Mode()
		// fmt.Println("Owner read:", fileMode&0400 != 0)
		// fmt.Println("Owner write:", fileMode&0200 != 0)
		// fmt.Println("Owner execute:", fileMode&0100 != 0)
		// fmt.Println("Group read:", fileMode&0040 != 0)
		// fmt.Println("Group write:", fileMode&0020 != 0)
		// fmt.Println("Group execute:", fileMode&0010 != 0)
		// fmt.Println("Others read:", fileMode&0004 != 0)
		// fmt.Println("Others write:", fileMode&0002 != 0)
		// fmt.Println("Others execute:", fileMode&0001 != 0)

		// decode file permission
		var permissionArr []string
		fileMode := finfo.Mode()

		if fileMode&0400 != 0 {
			permissionArr = append(permissionArr, "Read")
		}

		if fileMode&0200 != 0 {
			permissionArr = append(permissionArr, "Write")
		}

		if fileMode&0100 != 0 {
			permissionArr = append(permissionArr, "Execute")
		}

		permissionString := strings.Join(permissionArr, ", ")

		data = append(data, types.DirData{
			Name:         d.Name(),
			Size:         fSizeInKB,
			Type:         finfo.IsDir(),
			Path:         path + "\\" + d.Name(),
			Permission:   permissionString,
			DateModified: finfo.ModTime().Format("2006-01-02 15:04:05"),
		})

		// count total number of files and folders
		if finfo.IsDir() {
			totalFolders++
		} else {
			totalFiles++
		}

	}
	result := Result{
		TotalFiles:   totalFiles,
		TotalFolders: totalFolders,
		DirData:      data,
	}

	return result, nil
}

// CreateFileNFolder create file and folder
func CreateFileNFolder(rData types.ReqDirData) error {

	if rData.Type == "file" {
		slog.Info("Create file", "path", rData.CurrentDirPath+"\\"+rData.Name+rData.Extension)
		f, err := os.Create(rData.CurrentDirPath + "\\" + rData.Name + rData.Extension)
		if err != nil {
			return err
		}
		defer f.Close()

	} else {
		slog.Info("Create folder", "path", rData.CurrentDirPath+"\\"+rData.Name)
		err := os.Mkdir(rData.CurrentDirPath+"\\"+rData.Name, 0755)
		if err != nil {
			return err
		}
	}

	return nil
}

// DeleteFileNFolder delete file and folder
func DeleteFileNFolder(rData types.ReqDirData) error {

	err := os.RemoveAll(rData.CurrentDirPath)
	if err != nil {
		return err
	}

	return nil
}

// RenameFileNFolder rename file and folder
func RenameFileNFolder(rData types.RenameData) error {

	err := os.Rename(rData.OldName, rData.NewName)
	if err != nil {
		return err
	}

	return nil
}

// Read file content
func ReadFile(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	// contentWithBreaks := strings.ReplaceAll(string(content), "\n", "<br>")
	return string(content), nil
}
