package controllers

import (
	"github.com/bchanona/Api_chombi.git/src/users/application/usecases"
	"github.com/bchanona/Api_chombi.git/src/users/domain/entities/User/input"
	"github.com/gin-gonic/gin"
)

type RegisterUserController struct{
	useCase *usecases.RegisterUserUseCase
}


func NewRegisterUserController(useCase *usecases.RegisterUserUseCase) *RegisterUserController {
	return &RegisterUserController{useCase: useCase}
}

func (controller *RegisterUserController) Execute(ctx *gin.Context){
	var user input.UserRequest

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	
	if err := controller.useCase.Execute(user); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "User registered successfully"})

}
