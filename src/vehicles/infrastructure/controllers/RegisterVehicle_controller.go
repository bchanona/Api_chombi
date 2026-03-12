package controllers

import (
	"net/http"
	"github.com/bchanona/Api_chombi.git/src/vehicles/application/usecases"
	"github.com/bchanona/Api_chombi.git/src/vehicles/domain/entities/Vehicles/input"
	"github.com/gin-gonic/gin"
)

type RegisterVehicleController struct {
	useCase *usecases.RegisterVehicleUseCase
}

func NewRegisterVehicleController(useCase *usecases.RegisterVehicleUseCase) *RegisterVehicleController {
	return &RegisterVehicleController{
		useCase: useCase,
	}
}

func (controller *RegisterVehicleController) Execute(ctx *gin.Context) {

	var vehicleRequest input.VehicleRequest

	// Bind de datos del formulario
	if err := ctx.ShouldBind(&vehicleRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid vehicle data",
			"details": err.Error(),
		})
		return
	}

	// Obtener usuario desde JWT middleware
	userId, exist := ctx.Get("id")
	if !exist {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "User id not found in token",
		})
		return
	}

	vehicleRequest.UserId = userId.(string)

	// Obtener archivo
	file, header, err := ctx.Request.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Image file is required",
		})
		return
	}
	defer file.Close()

	// Ejecutar caso de uso
	err = controller.useCase.Execute(
		vehicleRequest,
		file,
		header.Filename,
	)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Vehicle registered successfully",
	})
}