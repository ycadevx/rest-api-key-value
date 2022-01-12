package main

import (
	"github/deliveryhero/source/api"
	"github/deliveryhero/source/log"
	. "github/deliveryhero/source/models"
	. "github/deliveryhero/source/storage"
)

func main() {

	TimeControl()
	CheckExistingData()

	log.Info.Printf("Server starting %s", API_PORT)

	api.StartAPI()

}
