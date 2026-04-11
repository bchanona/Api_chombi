package usecases

import (
	"github.com/bchanona/Api_chombi.git/src/vehicles/domain/repositories"
)

type DeleteVehicleUseCase struct {
	VehicleRepository repositories.IVehicleRepository
}

func NewDeleteVehicleUseCase(vehicleRepository repositories.IVehicleRepository) *DeleteVehicleUseCase {
	return &DeleteVehicleUseCase{VehicleRepository: vehicleRepository}
}

func (uc *DeleteVehicleUseCase) Execute(userId string, vehicleId string) error {
	err := uc.VehicleRepository.DeleteVehicle(userId, vehicleId)
	if err != nil {
		return err
	}
	return nil
}
