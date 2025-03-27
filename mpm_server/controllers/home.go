package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHome(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)

}

func GetSignup(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}
