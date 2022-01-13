package handler

import (
	"encoding/json"
	"fmt"
	"github/deliveryhero/source/log"
	_ "github/deliveryhero/source/log"
	. "github/deliveryhero/source/models"

	"net/http"
)

// Get All  Data
func GetAllData(w http.ResponseWriter, r *http.Request) {
	response := new(ApiResponse)
	defer RecoverPanic(r, response)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	switch r.Method {
	case http.MethodGet:
		w.WriteHeader(http.StatusOK)
		CreateResponse(w, &InMemoryDB)
		return

	default:
		err := http.StatusText(http.StatusMethodNotAllowed)
		response.Error = err
		w.WriteHeader(http.StatusMethodNotAllowed)
		CreateResponse(w, &response)
		return
	}
}

//Get Single Data
func GetData(w http.ResponseWriter, r *http.Request) {
	response := new(ApiResponse)
	defer RecoverPanic(r, response)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	switch r.Method {
	case http.MethodGet:
		if key, ok := r.URL.Query()["key"]; ok {
			if value, ok := InMemoryDB[key[0]]; ok {
				response.Result = value
				w.WriteHeader(http.StatusOK)
				CreateResponse(w, &response)
				return
			}
			w.WriteHeader(http.StatusNotFound)
			err := fmt.Sprintf(KeyNotFoundErrorResult, key[0])
			response.Error = err
			CreateResponse(w, &response)
			return
		}
		response.Error = KeyErrorResult
		w.WriteHeader(http.StatusBadRequest)
		CreateResponse(w, &response)
		return
	default:
		err := http.StatusText(http.StatusMethodNotAllowed)
		response.Error = err
		w.WriteHeader(http.StatusMethodNotAllowed)
		CreateResponse(w, &response)
		return

	}
}

//Add New Record
func AddData(w http.ResponseWriter, r *http.Request) {
	response := new(ApiResponse)
	defer RecoverPanic(r, response)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	switch r.Method {
	case http.MethodPost:
		var body User
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			//TO DO: YCA
			log.Error.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			CreateResponse(w, nil)
			return
		}

		if body.Key == "" {
			response.Error = KeyErrorResult
			w.WriteHeader(http.StatusBadRequest)
			CreateResponse(w, &response)
			return
		}

		if body.Value == "" {
			response.Error = ValueErrorResult
			w.WriteHeader(http.StatusBadRequest)
			CreateResponse(w, &response)
			return
		}

		StoreData(body.Key, body.Value)
		response.Result = fmt.Sprintf(SetResponseResult, body.Value, body.Key)
		w.WriteHeader(http.StatusCreated)
		CreateResponse(w, &response)
		return

	default:
		err := http.StatusText(http.StatusMethodNotAllowed)
		response.Error = err
		w.WriteHeader(http.StatusMethodNotAllowed)
		CreateResponse(w, &response)
		return
	}

}

//Delete All Data
func FlushData(w http.ResponseWriter, r *http.Request) {
	response := new(ApiResponse)
	defer RecoverPanic(r, response)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	switch r.Method {
	case http.MethodDelete:
		InMemoryDB = make(map[string]string)
		response.Result = FlushResponseResult
		w.WriteHeader(http.StatusNoContent)
		return

	default:
		err := http.StatusText(http.StatusMethodNotAllowed)
		response.Error = err
		w.WriteHeader(http.StatusMethodNotAllowed)
		CreateResponse(w, &response)
		return
	}
}
