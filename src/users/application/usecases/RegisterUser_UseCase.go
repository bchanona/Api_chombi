package usecases

import (
	"github.com/bchanona/Api_chombi.git/src/helper"
	"github.com/bchanona/Api_chombi.git/src/users/domain/entities/User/input"
	"github.com/bchanona/Api_chombi.git/src/users/domain/repositories"
)

type RegisterUserUseCase struct {
	db repositories.IUserRepository
}

func NewRegisterUserUseCase(db repositories.IUserRepository) *RegisterUserUseCase {
	return &RegisterUserUseCase{db: db}
}

func (useCase *RegisterUserUseCase) Execute(user input.UserRequest) error {
	//Genera uuid para id del usuario
	id_uuid := helper.GenerateUUID()
	user.ID = id_uuid

	//Encriptar contraseña
	encryptedPassword, err := helper.EncryptPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = encryptedPassword

	return useCase.db.RegisterUser(user)
} 