package auth

import (
	"net/http"

	"banter/handlers"

	"github.com/gin-gonic/gin"
)

var RouteGroupName = "/auth"

func Routes(router *gin.RouterGroup) {
	router.Handle(http.MethodPost, "/login", handlers.LoginHandler)
	router.Handle(http.MethodPost, "/register", handlers.RegisterHandler)

	return
}
