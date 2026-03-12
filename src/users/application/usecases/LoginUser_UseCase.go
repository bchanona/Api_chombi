package usecases

import (
	"github.com/bchanona/Api_chombi.git/src/helper"
	"github.com/bchanona/Api_chombi.git/src/users/domain/entities/User/input"
	"github.com/bchanona/Api_chombi.git/src/users/domain/entities/User/output"
	"github.com/bchanona/Api_chombi.git/src/users/domain/repositories"
)

type LoginUserUseCase struct {
	db repositories.IUserRepository
}

func NewLoginUserUseCase(db repositories.IUserRepository) *LoginUserUseCase {
	return &LoginUserUseCase{db: db}
}

func (useCase *LoginUserUseCase) Execute(loginReq input.UserLoginRequest) (output.UserLoginResponse, error) {
	email := loginReq.Email
	user, err := useCase.db.FindByEmail(email)
	if err != nil {
		return output.UserLoginResponse{}, err
	}
	//Si encuentra el usuario, se verifica la contraseña
	login := helper.ComparePassword(user.Password, loginReq.Password)
	if !login {
		return output.UserLoginResponse{}, err
	}
	//Generar token JWT
	token, err := helper.GenerateJWT(user.Id)
	if err != nil {
		return output.UserLoginResponse{}, err
	}
	user.Token = token

	return output.UserLoginResponse{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		RoleName:  user.RoleName,
		Token:     user.Token,
	}, nil

}
