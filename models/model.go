package model

import "github/deliveryhero/source/utils"

// Response object
type ApiResponse struct {
	Error  string `json:"error,omitempty"`
	Result string `json:"result,omitempty"`
}

// Request object
type User struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// Variable description
var (
	InMemoryDB             = make(map[string]string)
	KeyErrorResult         = "The 'key' is required."
	ValueErrorResult       = "The 'value' is required."
	KeyNotFoundErrorResult = "The key '%s' could'nt be found in data."
	SetResponseResult      = "The value '%s' is stored with the key '%s'."
	FlushResponseResult    = "All data has been deleted."
	EXPORT_FILE_PATH       = utils.GetEnv("EXPORT_FILE_PATH", "./tmp")
	API_PORT               = utils.GetEnv("API_PORT", "8085")
	RECORD_TIME            = utils.GetEnv("RECORD_TIME", "2")
)
