package v1

import (
	"banter/handlers"
	"banter/middlewares"
	"net/http"

	_ "banter/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var RouteGroupName = "/v1"

func UserRoutes(router *gin.RouterGroup) {
	router.Use(middlewares.JWTMiddleware())
	{
		// User related routes
		router.Handle(http.MethodGet, "/user/:id", handlers.GetUserDetailsHandler)
		router.Handle(http.MethodPatch, "/user/:id", handlers.UpdateUserDetailsHandler)

	}
}

func ConversationRoutes(router *gin.RouterGroup) {
	router.Use(middlewares.JWTMiddleware())
	{
		// User related routes
		router.Handle(http.MethodPost, "/conversation", handlers.StartConversationHandler)
		router.Handle(http.MethodGet, "/conversations/member/:user_id", handlers.GetConversationsHandler)
		router.Handle(http.MethodGet, "/conversation/:id", handlers.GetConversationHandler)
		router.Handle(http.MethodPost, "/conversation/:id/member/:user_id", handlers.AddMemberHandler)
		router.Handle(http.MethodDelete, "/conversation/:id/member/:user_id", handlers.RemoveMemberHandler)
		router.Handle(http.MethodDelete, "/conversation/:id", handlers.DeleteConversationHandler)

	}
}

func ApiDocRoutes(router *gin.RouterGroup) {
	router.Use()
	{
		// API documentation
		router.Handle(http.MethodGet, "/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}
