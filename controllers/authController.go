package controllers

import (
	"DMAPI/controllers/api"
	"DMAPI/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type _auth struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func auth(c *gin.Context) {
	var req _auth

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.ResponseError(err.Error(), 3))
		return
	}

	user := models.Auth(req.Login, req.Password)

	if user == nil {
		c.JSON(404, api.Response{Error: &api.Error{Message: "Неверный логин или пароль", Code: 404}})
		return
	}

	c.JSON(http.StatusOK, api.Response{Result: user})
}
