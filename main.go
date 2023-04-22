package main

import (
	"DMAPI/logger"
	"DMAPI/server"
)

func main() {
	logger.InfoNew("START!")
	server.Run()
}
