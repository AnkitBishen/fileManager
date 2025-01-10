package types

type PathT struct {
	Path string `json:"path"`
}

type DirData struct {
	Name         string `json:"name"`
	Type         bool   `json:"type"`
	Size         int64  `json:"size"`
	Permission   string `json:"permission"`
	Path         string `json:"path"`
	DateModified string `json:"dateModified"`
}

type ResponseErr struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

type ResponseSucc struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type ReqDirData struct {
	CurrentDirPath string
	Name           string
	Type           string
	Extension      string
}

type RenameData struct {
	OldName string
	NewName string
}
