package repository

import "github.com/bchanona/Api_chombi.git/src/users/domain/entities/User/input"

type IUserRepository interface {
	RegisterUser(input.UserRequest) error
}