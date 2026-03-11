package dependencies

import (
	"fmt"

	"github.com/bchanona/Api_chombi.git/src/helper/config"
	"github.com/bchanona/Api_chombi.git/src/users/application/usecases"
	"github.com/bchanona/Api_chombi.git/src/users/infrastructure"
	"github.com/bchanona/Api_chombi.git/src/users/infrastructure/controllers"
)

var (
	mySQL infrastructure.MySQL
)

func Init() {
	db, err := config.ConnMySQL()

	if err != nil {
		fmt.Println("Database connection error: ", err)
		return
	}

	mySQL = *infrastructure.NewMySQL(db)
}

func RegisterUserDependencies() *controllers.RegisterUserController {
	useCase := usecases.NewRegisterUserUseCase(&mySQL)
	return controllers.NewRegisterUserController(useCase)
}

func LoginUserDependencies() *controllers.LoginController {
	useCase := usecases.NewLoginUserUseCase(&mySQL)
	return controllers.NewLoginController(useCase)
}
