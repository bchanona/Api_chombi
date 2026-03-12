package repositories

import (
	"github.com/bchanona/Api_chombi.git/src/users/domain/entities/User/input"
	"github.com/bchanona/Api_chombi.git/src/users/domain/entities/User/output"
)

type IUserRepository interface {
	RegisterUser(input.UserRequest) error
	FindByEmail(email string) (output.UserLoginDBResponse, error)
}