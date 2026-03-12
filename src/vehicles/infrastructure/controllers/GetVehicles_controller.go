package controllers

import (
	"github.com/bchanona/Api_chombi.git/src/vehicles/application/usecases"
	"github.com/gin-gonic/gin"
)

type GetVehiclesController struct {
	useCase *usecases.GetVehiclesUseCase
}

func NewGetVehiclesController(useCase *usecases.GetVehiclesUseCase) *GetVehiclesController {
	return &GetVehiclesController{useCase: useCase}
}
func (controller *GetVehiclesController) Execute(ctx *gin.Context){
	userId, exists := ctx.Get("id")
	if !exists {
		ctx.JSON(401, gin.H{"error": "User ID not found in token"})
		return
	}

	vehicles, err := controller.useCase.Execute(userId.(string))
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to retrieve vehicles"})
		return
	}

	ctx.JSON(200, gin.H{"vehicles": vehicles})
}