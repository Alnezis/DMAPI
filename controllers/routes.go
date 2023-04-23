package controllers

import (
	"DMAPI/controllers/contractors"
	"DMAPI/controllers/deals"
	"DMAPI/controllers/documents"
	"DMAPI/controllers/files"
	"DMAPI/controllers/messages"
	"DMAPI/controllers/photos"
	"DMAPI/controllers/static"
	gs "github.com/gin-contrib/static"
)

func (server *Server) initializeRoutes() {

	server.Router.POST("/auth", auth)

	server.Router.GET("/files/:filename", files.Files)

	server.Router.GET("/documents/list/:what", documents.GetDocuments)
	server.Router.POST("/documents/setStatus", documents.SetStatusDoc)

	server.Router.POST("/files/upload", files.HandlePost)

	server.Router.GET("/contractor/list", contractors.GetContractorsInfo)

	server.Router.GET("/deals/list", deals.GetDeals)
	server.Router.POST("/deals/create", deals.CreateDeal)
	server.Router.GET("/deals/get", deals.GetDeal)
	server.Router.POST("/deals/setStatus", deals.SetStatusDeal)

	server.Router.POST("/photos/attach", photos.AddPhoto)

	server.Router.POST("/messages/send", messages.Send)
	server.Router.GET("/messages/list", messages.MessagesList)

	server.Router.GET("/static/statuses", static.GetStatuses)

	server.Router.Use(gs.Serve("/", gs.LocalFile("/var/www/html", false)))

	///var/www/html
}
