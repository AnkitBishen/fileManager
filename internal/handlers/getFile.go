package getFile

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/AnkitBishen/fileManagerApp/internal/funcs"
	"github.com/AnkitBishen/fileManagerApp/internal/response"
	"github.com/AnkitBishen/fileManagerApp/internal/types"
)

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
