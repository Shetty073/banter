package auth

import (
	"net/http"

	"banter/controllers"

	"github.com/gin-gonic/gin"
)

var RouteGroupName = "/auth"

func Routes(router *gin.RouterGroup) {
	router.Handle(http.MethodPost, "/login", controllers.LoginController)
	router.Handle(http.MethodPost, "/register", controllers.RegisterController)

	return
}
