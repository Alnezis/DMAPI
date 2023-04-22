package server

import (
	"DMAPI/controllers"
)

var server = controllers.Server{}

func Run() {
	server.Initialize()
	server.Run()

}
