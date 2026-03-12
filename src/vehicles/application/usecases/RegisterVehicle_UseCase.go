package usecases

import (
	"mime/multipart"

	"github.com/bchanona/Api_chombi.git/src/helper"
	"github.com/bchanona/Api_chombi.git/src/vehicles/domain/entities/Vehicles/input"
	"github.com/bchanona/Api_chombi.git/src/vehicles/domain/ports"
	"github.com/bchanona/Api_chombi.git/src/vehicles/domain/repositories"
)

type RegisterVehicleUseCase struct {
	db repositories.IVehicleRepository
	imagesStorage ports.ImageStorage
}

func NewRegisterVehicleUseCase(db repositories.IVehicleRepository, imagesStorage ports.ImageStorage) *RegisterVehicleUseCase {
	return &RegisterVehicleUseCase{db: db, imagesStorage: imagesStorage}
}

func (useCase *RegisterVehicleUseCase) Execute(vehicleRequest input.VehicleRequest, file multipart.File, fileName string) error {
	//Generar uuid para id del vehículo
	id__uuid := helper.GenerateUUID()
	vehicleRequest.Id = id__uuid
	//Definir pasajeros en 0 por defecto
	vehicleRequest.NumberOfPassengers = 0
	vehicleRequest.Shift = 1
	cloudinaryImageURL, err := useCase.imagesStorage.UploadImage(file, fileName)

	if err != nil {
		return err
	}
	vehicleRequest.ImageURL = cloudinaryImageURL
	return useCase.db.RegisterVehicle(vehicleRequest)
}