package handler

import (
	"encoding/json"
	"github/deliveryhero/source/log"
	. "github/deliveryhero/source/models"
	"net/http"

	"github.com/fatih/color"
)

func CreateResponse(w http.ResponseWriter, response interface{}) {
	if err := json.NewEncoder(w).Encode(response); err != nil {
		w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func RecoverPanic(r *http.Request, response *ApiResponse) {
	if rec := recover(); rec != nil {
		log.Error.Printf("[RECOVERED PANIC] %+v", rec)
	}
	log.Info.Printf("%s  Response: {%+v}]", color.CyanString(r.Method), response)
}

//Allows data to be written to InMemoryDB.
func StoreData(key, value string) {
	InMemoryDB[key] = value
}
