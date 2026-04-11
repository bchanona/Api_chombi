package dependencies

import (
	"github.com/bchanona/Api_chombi.git/src/helper/config"
	"github.com/bchanona/Api_chombi.git/src/helper/services"
	"github.com/bchanona/Api_chombi.git/src/vehicles/application/usecases"
	"github.com/bchanona/Api_chombi.git/src/vehicles/infrastructure"
	"github.com/bchanona/Api_chombi.git/src/vehicles/infrastructure/controllers"
)

var (
	mySQL             infrastructure.MySQL
	cloudinaryService services.CloudinaryService
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

func GetVehiclesDependencies() *controllers.GetVehiclesController {
	useCase := usecases.NewGetVehiclesUseCase(&mySQL)
	return controllers.NewGetVehiclesController(useCase)
}

func RegisterVehicleShiftHistoryDependencies() *controllers.RegisterVehicleHistoryController {
	useCase := usecases.NewRegisterVehicleHistoryUseCase(&mySQL)
	return controllers.NewRegisterVehicleHistoryController(useCase)
}
func UploadPdfDependencies() *controllers.UploadPdfController {
	useCase := usecases.NewUploadPdfUseCase(&cloudinaryService)
	return controllers.NewUploadPdfController(useCase)
}
func UpdateVehicleDependencies() *controllers.UpdateVehicleController {
	useCase := usecases.NewUpdateVehicleUseCase(&mySQL)
	return controllers.NewUpdateVehicleController(useCase)
}

func DeleteVehicleDependencies() *controllers.DeleteVehicleController {
	useCase := usecases.NewDeleteVehicleUseCase(&mySQL)
	return controllers.NewDeleteVehicleController(useCase)
}

func GetVehicleByUnitNumberDependencies() *controllers.GetVehicleByUnitNumberController {
	useCase := usecases.NewGetVehicleByUnitNumberUseCase(&mySQL)
	return controllers.NewGetVehicleByUnitNumberController(useCase)
}

func GetVehicleHistoryByDateDependencies() *controllers.GetVehicleHistoryByDateController {
	useCase := usecases.NewGetVehicleHistoryByDateUseCase(&mySQL)
	return controllers.NewGetVehicleHistoryByDateController(useCase)
}
