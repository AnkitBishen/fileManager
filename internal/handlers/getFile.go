package getFile

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/AnkitBishen/fileManagerApp/internal/funcs"
	"github.com/AnkitBishen/fileManagerApp/internal/response"
	"github.com/AnkitBishen/fileManagerApp/internal/types"
)

// Handler func for get List of files and folders
func List() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("handler Get list of files and folders", "path", r.URL.Path)

		var path types.PathT

		err := json.NewDecoder(r.Body).Decode(&path)
		if err != nil {
			fmt.Println(err.Error())
			response.WriteJson(w, http.StatusBadRequest, response.Error(err))
			return
		}

		// get list of folders and files
		data, err := funcs.GetFolderAndFile(path.Path)
		if err != nil {
			fmt.Println(err.Error())
			response.WriteJson(w, http.StatusInternalServerError, response.Error(err))
			return
		}

		response.WriteJson(w, http.StatusOK, response.Success(data))
	}
}

// Handler func for create folder and files
func CreateDir() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var rData types.ReqDirData
		err := json.NewDecoder(r.Body).Decode(&rData)
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.Error(err))
			return
		}

		err = funcs.CreateFileNFolder(rData)
		if err != nil {
			if os.IsExist(err) {
				response.WriteJson(w, http.StatusForbidden, response.Error(err))
				return
			}
			response.WriteJson(w, http.StatusInternalServerError, response.Error(err))
			return
		}

		response.WriteJson(w, http.StatusOK, response.Success(rData.Type+" created successfully."))
	}
}

// Handler func for delete folder and files
func Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var rData types.ReqDirData

		err := json.NewDecoder(r.Body).Decode(&rData)
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.Error(err))
			return
		}

		err = funcs.DeleteFileNFolder(rData)
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.Error(err))
			return
		}

		response.WriteJson(w, http.StatusOK, response.Success(rData.CurrentDirPath+" deleted successfully."))
	}
}

// Handler func for Rename folder and files
func Rename() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var rData types.RenameData

		err := json.NewDecoder(r.Body).Decode(&rData)
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.Error(err))
			return
		}

		err = funcs.RenameFileNFolder(rData)
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.Error(err))
			return
		}

		response.WriteJson(w, http.StatusOK, response.Success(rData.NewName+" changed successfully."))
	}
}
