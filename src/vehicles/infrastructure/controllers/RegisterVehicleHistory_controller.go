package controllers

import (
	"github.com/bchanona/Api_chombi.git/src/vehicles/application/usecases"
	"github.com/gin-gonic/gin"
)

type RegisterVehicleHistoryController struct {
	useCase *usecases.RegisterVehicleHistoryUseCase
}

func NewRegisterVehicleHistoryController(registerVehicleHistoryUseCase *usecases.RegisterVehicleHistoryUseCase) *RegisterVehicleHistoryController {
	return &RegisterVehicleHistoryController{useCase: registerVehicleHistoryUseCase}
}

func (controller *RegisterVehicleHistoryController) Execute(ctx *gin.Context){
	err := controller.useCase.Execute()
    
    if err != nil {
        ctx.JSON(500, gin.H{
            "error":   "Failed to register vehicle history",
            "details": err.Error(),
        })
        return
    }

    // Respuesta de éxito
    ctx.JSON(200, gin.H{
        "message": "Vehicle shift history registered successfully",
    })
}