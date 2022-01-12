package api

import (
	"fmt"

	"net/http"

	. "github/deliveryhero/source/controller"
	"github/deliveryhero/source/log"
	. "github/deliveryhero/source/models"
)

func StartAPI() {
	//endpoints
	mux := http.NewServeMux()

	mux.HandleFunc("/", GetAllData)
	mux.HandleFunc("/get", GetData)
	mux.HandleFunc("/set", AddData)
	mux.HandleFunc("/flush", FlushData)

	log.Fatal.Println(http.ListenAndServe(fmt.Sprintf(":%s", API_PORT), mux))
}
