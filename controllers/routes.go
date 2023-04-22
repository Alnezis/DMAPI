package controllers

import (
	"DMAPI/controllers/contractors"
	"DMAPI/controllers/deals"
	"DMAPI/controllers/documents"
	"DMAPI/controllers/files"
	"DMAPI/controllers/messages"
	"DMAPI/controllers/photos"
)

func (server *Server) initializeRoutes() {
	server.Router.POST("/auth", auth)
	server.Router.GET("/files/:filename", files.Files)

	server.Router.GET("/documents/list/:what", documents.GetDocuments)

	server.Router.POST("/files/upload", files.HandlePost)

	server.Router.GET("/contractor/list", contractors.GetContractorsInfo)
	server.Router.GET("/deals/list", deals.GetDeals)
	server.Router.POST("/deals/create", deals.CreateDeal)
	server.Router.GET("/deals/get", deals.GetDeal)

	server.Router.POST("/photos/attach", photos.AddPhoto)

	server.Router.POST("/messages/send", messages.Send)
	server.Router.GET("/messages/list", messages.MessagesList)

	//server.Router.Any("/getCroppedPlan", getCroppedPlan)

}
