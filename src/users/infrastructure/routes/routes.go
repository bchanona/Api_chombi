package routes

import (
	"github.com/bchanona/Api_chombi.git/src/users/infrastructure/dependencies"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.RouterGroup){
	routes := router.Group("/auth")

	registerUserController := dependencies.RegisterUserDependencies().Execute

	routes.POST("/register",registerUserController)
	

}