package controllers

import (
	"github.com/bchanona/Api_chombi.git/src/vehicles/application/usecases"
	"github.com/gin-gonic/gin"
)

type GetVehicleHistoryByDateController struct {
	GetVehicleHistoryByDateUseCase *usecases.GetVehicleHistoryByDateUseCase
}

func NewGetVehicleHistoryByDateController(getVehicleHistoryByDateUseCase *usecases.GetVehicleHistoryByDateUseCase) *GetVehicleHistoryByDateController {
	return &GetVehicleHistoryByDateController{GetVehicleHistoryByDateUseCase: getVehicleHistoryByDateUseCase}
}

func (ctrl *GetVehicleHistoryByDateController) Execute(ctx *gin.Context) {

	userdId, exists := ctx.Get("id")
	if !exists {
		ctx.JSON(401, gin.H{"error": "User ID not found in token"})
		return
	}

	date := ctx.Param("date")
	if date == "" {
		ctx.JSON(400, gin.H{"error": "date parameter is required (format: YYYY-MM-DD)"})
		return
	}

	history, err := ctrl.GetVehicleHistoryByDateUseCase.Execute(userdId.(string), date)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, history)

}
