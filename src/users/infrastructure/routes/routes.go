package routes

import (
	"github.com/bchanona/Api_chombi.git/src/users/infrastructure/dependencies"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.RouterGroup) {
	routes := router.Group("/auth")

	registerUserController := dependencies.RegisterUserDependencies().Execute
	loginUserController := dependencies.LoginUserDependencies().Execute

	routes.POST("/register", registerUserController)
	routes.POST("/login", loginUserController)

}
