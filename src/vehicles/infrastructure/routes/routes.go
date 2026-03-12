package routes

import (
	"github.com/bchanona/Api_chombi.git/src/helper/middlewares"
	"github.com/bchanona/Api_chombi.git/src/vehicles/infrastructure/dependencies"
	"github.com/gin-gonic/gin"
)

func VehicleRoutes(router *gin.RouterGroup) {

	routes := router.Group("/vehicles")

	registerVehicleController := dependencies.RegisterVehicleDependencies()
	getVehicleHistoryController := dependencies.GetVehiclesByUserIdDependencies()

	routes.POST(
		"/",
		middlewares.AuthMiddleware(),
		registerVehicleController.Execute,
	)
	routes.GET(
		"/history",
		middlewares.AuthMiddleware(),
		getVehicleHistoryController.Execute,
	)
	routes.GET(
		"/",
		middlewares.AuthMiddleware(),
		dependencies.GetVehiclesDependencies().Execute,
	)

}