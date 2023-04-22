package controllers

import (
	"DMAPI/app"
	"DMAPI/logger"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"os"
	"time"
)

type Server struct {
	Router *gin.Engine
}

func (server *Server) Initialize() {

	gin.SetMode(gin.ReleaseMode)

	server.Router = gin.Default()

	server.Router.SetTrustedProxies(nil)

	cs := cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	})
	server.Router.Use(cs)

	//server.Router.Use(cors.Default())

	server.initializeRoutes()
}

func (server *Server) Run() {

	cert := app.CFG.SslPath + "fullchain.pem"
	key := app.CFG.SslPath + "privkey.pem"

	if _, err := os.Stat(cert); err != nil {

		if os.IsNotExist(err) {
			logger.Info.Println("Good start: SSL NO")
			err := server.Router.Run(fmt.Sprintf(":%s", app.CFG.Port))
			if err != nil {
				logger.Error.Println(err)
			}
		}
		logger.Error.Println(err)
	} else {
		logger.Info.Println("Good start: SSL YES")
		err := server.Router.RunTLS(fmt.Sprintf(":%s", app.CFG.Port), cert, key)
		if err != nil {
			logger.Error.Println(err)
		}
	}
}
