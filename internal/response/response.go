package response

import (
	"encoding/json"
	"net/http"

	"github.com/AnkitBishen/fileManagerApp/internal/types"
)

func WriteJson(w http.ResponseWriter, statusCode int, data interface{}) error {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	return json.NewEncoder(w).Encode(&data)
}

func Error(err error) types.ResponseErr {

	return types.ResponseErr{
		Status: "failed",
		Error:  err.Error(),
	}
}

func Success(data interface{}) types.ResponseSucc {
	return types.ResponseSucc{
		Status: "success",
		Data:   data,
	}
}
