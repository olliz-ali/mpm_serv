package mpmserverfrontend

import (
	"github.com/gin-gonic/gin"
)

func LoadHTMLs(r *gin.Engine) {
	r.LoadHTMLFiles("mpm_server_frontend/templates/index.html", "mpm_server_frontend/templates/signup.html")
}

func LoadCSSs(r *gin.Engine) {
	r.Static("static/css", "mpm_server_frontend/static/css")
	r.Static("static/js", "mpm_server_frontend/static/js")
	r.Static("assets", "mpm_server_frontend/assets")
}
