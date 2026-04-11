package controllers

import (
	"strconv"

	"github.com/bchanona/Api_chombi.git/src/vehicles/application/usecases"
	"github.com/gin-gonic/gin"
)

type GetVehicleByUnitNumberController struct {
	GetVehicleByUnitNumberUseCase *usecases.GetVehicleByUnitNumberUseCase
}

func NewGetVehicleByUnitNumberController(getVehicleByUnitNumberUseCase *usecases.GetVehicleByUnitNumberUseCase) *GetVehicleByUnitNumberController {
	return &GetVehicleByUnitNumberController{GetVehicleByUnitNumberUseCase: getVehicleByUnitNumberUseCase}
}

func (ctrl *GetVehicleByUnitNumberController) Execute(ctx *gin.Context) {

	userdId, exists := ctx.Get("id")
	if !exists {
		ctx.JSON(401, gin.H{"error": "User ID not found in token"})
		return
	}

	unitNumberStr := ctx.Param("unitNumber")
	if unitNumberStr == "" {
		ctx.JSON(400, gin.H{"error": "unitNumber parameter is required"})
		return
	}

	unitNumber, err := strconv.Atoi(unitNumberStr)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "unitNumber must be a valid integer"})
		return
	}

	vehicle, err := ctrl.GetVehicleByUnitNumberUseCase.Execute(userdId.(string), unitNumber)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, vehicle)

}
