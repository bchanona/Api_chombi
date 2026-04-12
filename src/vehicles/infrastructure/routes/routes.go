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
	routes.POST(
		"/history",
		middlewares.AuthMiddleware(),
		dependencies.RegisterVehicleShiftHistoryDependencies().Execute,
	)
	routes.POST(
		"/upload-pdf",
		middlewares.AuthMiddleware(),
		dependencies.UploadPdfDependencies().Execute,
	)
	routes.GET(
		"/pdf-urls",
		middlewares.AuthMiddleware(),
		dependencies.GetPdfDependencies().Execute,
	)
	routes.PUT(
		"/:vehicleId",
		middlewares.AuthMiddleware(),
		dependencies.UpdateVehicleDependencies().Execute,
	)
	routes.DELETE(
		"/:vehicleId",
		middlewares.AuthMiddleware(),
		dependencies.DeleteVehicleDependencies().Execute,
	)
	routes.GET(
		"/by-unit/:unitNumber",
		middlewares.AuthMiddleware(),
		dependencies.GetVehicleByUnitNumberDependencies().Execute,
	)
	routes.GET(
		"/history/by-date/:date",
		middlewares.AuthMiddleware(),
		dependencies.GetVehicleHistoryByDateDependencies().Execute,
	)

}
