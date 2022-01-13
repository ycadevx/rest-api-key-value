package api

import (
	"fmt"
	. "github/deliveryhero/source/controller"
	"github/deliveryhero/source/log"
	. "github/deliveryhero/source/models"
	"net/http"
)

func StartAPI() {
	//endpoints
	mux := http.NewServeMux()

	mux.HandleFunc("/api/get-all", GetAllData)
	mux.HandleFunc("/api/get", GetData)
	mux.HandleFunc("/api/set", AddData)
	mux.HandleFunc("/api/flush", FlushData)

	log.Fatal.Println(http.ListenAndServe(fmt.Sprintf(":%s", API_PORT), mux))
}
