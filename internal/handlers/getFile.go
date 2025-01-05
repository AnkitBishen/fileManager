package getFile

import (
	"encoding/json"
	"net/http"

	"github.com/AnkitBishen/fileManagerApp/internal/response"
)

type pathT struct {
	Path string `json:"path"`
}

func List() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var path pathT

		err := json.NewDecoder(r.Body).Decode(&path)
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, err)
			return
		}

		response.WriteJson(w, http.StatusOK, path)
	}
}
