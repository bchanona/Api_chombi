package routes

import (
	"github.com/bchanona/Api_chombi.git/src/helper/middlewares"
	"github.com/bchanona/Api_chombi.git/src/vehicles/infrastructure/dependencies"
	"github.com/gin-gonic/gin"
)

func VehicleRoutes(router *gin.RouterGroup) {

	routes := router.Group("/vehicles")

	registerVehicleController := dependencies.RegisterVehicleDependencies()

	routes.POST(
		"/",
		middlewares.AuthMiddleware(),
		registerVehicleController.Execute,
	)
}