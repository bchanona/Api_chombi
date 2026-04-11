package usecases

import (
	"github.com/bchanona/Api_chombi.git/src/vehicles/domain/entities/Vehicles/input"
	"github.com/bchanona/Api_chombi.git/src/vehicles/domain/repositories"
)


type UpdateVehicleUseCase struct {
	VehicleRepository repositories.IVehicleRepository
}

func NewUpdateVehicleUseCase(vehicleRepository repositories.IVehicleRepository) *UpdateVehicleUseCase {
	return &UpdateVehicleUseCase{VehicleRepository: vehicleRepository}
}

func (uc *UpdateVehicleUseCase) Execute(userId string, vehicleId string, vehicleData input.UpdateVehicleRequest) error {
	err := uc.VehicleRepository.UpdateVehicle(userId, vehicleId, vehicleData)
	if err != nil {
		return err
	}
	return nil
}