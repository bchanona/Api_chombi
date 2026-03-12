package dependencies

import (
	"github.com/bchanona/Api_chombi.git/src/helper/config"
	"github.com/bchanona/Api_chombi.git/src/helper/services"
	"github.com/bchanona/Api_chombi.git/src/vehicles/application/usecases"
	"github.com/bchanona/Api_chombi.git/src/vehicles/infrastructure"
	"github.com/bchanona/Api_chombi.git/src/vehicles/infrastructure/controllers"
)

var (
	mySQL              infrastructure.MySQL
	cloudinaryService  services.CloudinaryService
)

func Init() {

	db, err := config.ConnMySQL()
	if err != nil {
		panic("database connection failed: " + err.Error())
	}

	mySQL = *infrastructure.NewMySQL(db)

	// inicializar cloudinary una vez
	config.InitCloudinary()

	cloudinaryService = services.CloudinaryService{}
}

func RegisterVehicleDependencies() *controllers.RegisterVehicleController {

	useCase := usecases.NewRegisterVehicleUseCase(
		&mySQL,
		&cloudinaryService,
	)

	return controllers.NewRegisterVehicleController(useCase)
}

func GetVehiclesByUserIdDependencies() *controllers.GetVehicleHistoryController {
	useCase := usecases.NewGetVehicleHistoryUseCase(&mySQL)
	return controllers.NewGetVehicleHistoryController(useCase)
}