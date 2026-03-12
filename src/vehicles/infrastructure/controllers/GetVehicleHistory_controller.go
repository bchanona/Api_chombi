package controllers

import (
	"github.com/bchanona/Api_chombi.git/src/vehicles/application/usecases"
	"github.com/gin-gonic/gin"
)

type GetVehicleHistoryController struct {
	useCase *usecases.GetVehicleHistoryUseCase
}

func NewGetVehicleHistoryController(useCase *usecases.GetVehicleHistoryUseCase) *GetVehicleHistoryController {
	return &GetVehicleHistoryController{useCase: useCase}
}
func (controller *GetVehicleHistoryController) Execute(ctx *gin.Context){
	userId, exists := ctx.Get("id")
	if !exists {
		ctx.JSON(400, gin.H{"error": "User ID is required"})
		return
	}


	vehicleHistory, err := controller.useCase.Execute(userId.( string))
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if len(vehicleHistory) == 0 {
		ctx.JSON(404, gin.H{"message": "No vehicle history found for this user"})
		return
	}
	ctx.JSON(200, gin.H{"vehicle_history": vehicleHistory})
}