package controllers

import (
	"github.com/bchanona/Api_chombi.git/src/users/application/usecases"
	"github.com/bchanona/Api_chombi.git/src/users/domain/entities/User/input"
	"github.com/gin-gonic/gin"
)

type LoginController struct {
	useCase *usecases.LoginUserUseCase
}

func NewLoginController(useCase *usecases.LoginUserUseCase) *LoginController {
	return &LoginController{useCase: useCase}
}

func (controller *LoginController) Execute(ctx *gin.Context) {
	var loginReq input.UserLoginRequest

	if err := ctx.ShouldBindJSON(&loginReq); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	user, err := controller.useCase.Execute(loginReq)
	if err != nil {
		ctx.JSON(401, gin.H{"error": "Invalid email or password"})
		return
	}

	ctx.JSON(200, user)

}
