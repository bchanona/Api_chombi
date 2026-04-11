package controllers

import (
	"github.com/bchanona/Api_chombi.git/src/vehicles/application/usecases"
	"github.com/bchanona/Api_chombi.git/src/vehicles/domain/entities/Vehicles/input"
	"github.com/gin-gonic/gin"
)

type UpdateVehicleController struct {
	UpdateVehicleUseCase *usecases.UpdateVehicleUseCase
}

func NewUpdateVehicleController(updateVehicleUseCase *usecases.UpdateVehicleUseCase) *UpdateVehicleController {
	return &UpdateVehicleController{UpdateVehicleUseCase: updateVehicleUseCase}
}

func (ctrl *UpdateVehicleController) Execute(ctx *gin.Context) {

	userdId, exists := ctx.Get("id")
	if !exists {
		ctx.JSON(401, gin.H{"error": "User ID not found in token"})
		return
	}

	var vehicleData input.UpdateVehicleRequest
	if err := ctx.ShouldBind(&vehicleData); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	vehicleId := ctx.Param("vehicleId")
	if vehicleId == "" {
		ctx.JSON(400, gin.H{"error": "Vehicle ID is required"})
		return
	}

	err := ctrl.UpdateVehicleUseCase.Execute(userdId.(string), vehicleId, vehicleData)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "Vehicle updated successfully"})

}
