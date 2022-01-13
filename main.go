package main

/*
created by yusuf cakir
e-mail : yusufckr6194@gmail.com
*/

import (
	"github/deliveryhero/source/api"
	"github/deliveryhero/source/log"
	. "github/deliveryhero/source/storage"
)

func main() {

	TimeControl()
	CheckData()

	log.Info.Println("Server starting")
	api.StartAPI()
}
