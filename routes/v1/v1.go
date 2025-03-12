package v1

import (
	"net/http"
	"banter/middlewares"
	"banter/responses"

	_ "banter/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var RouteGroupName = "/v1"

// Placeholder is a test controller
// @Summary Placeholder dummy
// @Description This is a dummp placeholder
// @Tags Placeholder
// @Accept json
// @Produce json
// @Param request body schemas.LoginSchema true "Placeholder sample"
// @Success 200 {object} responses.SuccessBody
// @Failure 400 {object} responses.FailureBody
// @Failure 401 {object} responses.FailureBody
// @Failure 500 {object} responses.FailureBody
// @Router /placeholder [post]
// @Security AuthorizationToken
func PlaceHolder(c *gin.Context) {
	responses.Ok(c, gin.H{"message": "Hello, world!"})
}

func UserRoutes(router *gin.RouterGroup) {
	router.Use(middlewares.JWTMiddleware())
	{
		// User related routes
		router.Handle(http.MethodPost, "/user/create", PlaceHolder)
		router.Handle(http.MethodPost, "/user/read", PlaceHolder)
		router.Handle(http.MethodPost, "/user/read/:id", PlaceHolder)
		router.Handle(http.MethodPost, "/user/update/:id", PlaceHolder)
		router.Handle(http.MethodPost, "/user/delete/:id", PlaceHolder)
	}
}

func ApiDocRoutes(router *gin.RouterGroup) {
	router.Use()
	{
		// API documentation
		router.Handle(http.MethodGet, "/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}
